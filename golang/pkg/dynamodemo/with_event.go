package dynamodemo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var tableNameMap = map[string]string{
	"post-office": "post_office_event",
	"push":        "post_office_push_event",
	"email":       "post_office_email_event",
}

func InsertEvent(
	ctx context.Context,
	key string,
	eventID string,
	stage string,
	data []byte,
) (err error) {
	realEventID := fmt.Sprintf("%s:%s", stage, eventID)
	item := map[string]*dynamodb.AttributeValue{
		"Owner":     new(dynamodb.AttributeValue).SetS(key),
		"EventID":   new(dynamodb.AttributeValue).SetS(realEventID),
		"CreatedAt": new(dynamodb.AttributeValue).SetS(getCurrentTime()),
		"TTL":       new(dynamodb.AttributeValue).SetN(fmt.Sprintf("%d", time.Now().Add(time.Hour*24*365*60).Unix())),
	}
	if stage == "post-office" {
		item["Raw"] = new(dynamodb.AttributeValue).SetB(data)
	}
	tableName = tableNameMap[stage]

	resp, err := ddbClient.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(#EventID)"),
		ExpressionAttributeNames: map[string]*string{
			"#EventID": aws.String("EventID"),
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
	})
	if awsErr, ok := err.(awserr.Error); ok {
		log.Printf("err: %s, code: %s, msg: %s", awsErr.Code(), awsErr.Code(), awsErr.Message())
		if awsErr.Code() == "ConditionalCheckFailedException" {
			log.Printf("negligible err: %+v", err)
			err = nil
		}
	}
	if err != nil {
		log.Printf("Failed to truckate list: %+v", err)
	}
	fmt.Printf("trucate resp: %s\n", resp.GoString())
	return
}

func getCurrentTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func QueryEvent(ctx context.Context, key string, stage string, n int) ([]map[string]*dynamodb.AttributeValue, error) {
	tableName := tableNameMap[stage]
	resp, err := ddbClient.QueryWithContext(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(idxCreatedAt),
		KeyConditionExpression: aws.String("#HashKey = :Owner"),
		ScanIndexForward:       aws.Bool(false),
		Limit:                  aws.Int64(int64(n)),
		ExpressionAttributeNames: map[string]*string{
			"#HashKey": aws.String("Owner"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Owner": new(dynamodb.AttributeValue).SetS(key),
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
	})
	log.Printf("query resp: %s", resp.GoString())
	if awsErr, ok := err.(awserr.Error); ok {
		log.Printf("err: %s, code: %s, msg: %s", awsErr.Code(), awsErr.Code(), awsErr.Message())
		if awsErr.Code() == "ConditionalCheckFailedException" {
			log.Printf("negligible err: %+v", err)
			err = nil
		}
	}
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}
