package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	sinkQueueURL = "https://us-west-2.queue.amazonaws.com/097605708335/fanout-sink"

	// DynamoDB schema
	tableName = "CommEvent"

	// Hash Key
	hashKeyOwner = "Owner"
	// Range Key
	rangeKeyEventID = "EventID"

	colCreatedAt = "CreatedAt"
	colTTL       = "TTL"

	colPoBox      = "PoBox"
	colCampaignID = "CampaignID"
)

var (
	eventCount uint64
	commCount  uint64
)

func inc(name string) {
	uptr := &eventCount
	if name == "comm-cnt" {
		uptr = &commCount
	}
	n := atomic.AddUint64(uptr, 1)
	if n%300 == 0 {
		log.Printf("[INFO][%s] %d events processed", name, n)
	}
}

type CommEvent struct {
	Owner   string `json:"owner,omitempty"`
	EventID string `json:"event_id,omitempty"`
	PoBox   string `json:"po_box,omitempty"`

	CampaignID string `json:"campaign_id,omitempty"`
}

func checkIdempotency(ctx context.Context, commEvent *CommEvent) (exists bool, err error) {
	now := time.Now().UTC()

	if isDebugMode {
		log.Printf("[DEBUG] in local test, so skipping checking idempotency")
		return false, nil
	}

	ddbClient := getDdbClient()
	_, err = ddbClient.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			hashKeyOwner:    val().SetS(fmt.Sprintf("test:fanout:" + commEvent.Owner)),
			rangeKeyEventID: val().SetS(commEvent.EventID),

			colCreatedAt: val().SetS(now.Format(time.RFC3339)),
			colTTL:       val().SetN(ttlVal(now.Add(30 * 24 * time.Hour))),

			colPoBox:      val().SetS(commEvent.PoBox),
			colCampaignID: val().SetS(commEvent.CampaignID),
		},
		ConditionExpression: aws.String("attribute_not_exists(#hashKeyOwner) AND attribute_not_exists(#rangeKeyEventID)"),
		ExpressionAttributeNames: map[string]*string{
			"#hashKeyOwner":    aws.String(hashKeyOwner),
			"#rangeKeyEventID": aws.String(rangeKeyEventID),
		},
	})
	if awsErr, ok := err.(awserr.Error); ok {
		// skip certain errors.
		if awsErr.Code() == "ConditionalCheckFailedException" {
			// it's idempotent, so we can skip it.
			exists = true
			err = nil
		}
	}
	return
}

func sendToSinkSqs(ctx context.Context, evt *CommEvent) error {
	msg, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	if isDebugMode {
		log.Printf("[DEBUG] in local test, so skipping sending to sink sqs")
		log.Printf("[DEBUG] sendToSinkSqs: %s", string(msg))
		return nil
	}

	sqsClient := getSQSClient()
	_, err = sqsClient.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(msg)),
		QueueUrl:    aws.String(sinkQueueURL),
	})
	if err != nil {
		log.Printf("[ERROR] failed to send msg for source id: %s, err: %+v\n", string(msg), err)
		return err
	}
	return nil
}

func ttlVal(t time.Time) string {
	return fmt.Sprintf("%d", t.Unix())
}
