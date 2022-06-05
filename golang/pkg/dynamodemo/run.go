package dynamodemo

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ColName string

var (
	ddbClient *dynamodb.DynamoDB

	tableName             = "atc"
	sortKeyV1Meta ColName = "history"

	colEmailHistory ColName = "EmailHistory"
	colPushHistory  ColName = "PushHistory"

	tableNameV2 = "po_event_v2"

	tableNameForPostOfficeEvent = "post_office_event"
	tableNameForPushEvent       = "post_office_push_event"
	tableNameForEmailEvent      = "post_office_email_event"
	idxCreatedAt                = "IdxCreatedAt"
)

func Run() (err error) {
	ctx := context.Background()

	eventID := fmt.Sprintf("event-id-%d", time.Now().Unix())
	value := strings.Repeat(fmt.Sprintf("%d:%s", time.Now().Unix(), eventID), 10)

	awscfg := InitAws()
	sess := session.Must(session.NewSession(awscfg))
	ddbClient = dynamodb.New(sess)

	key := "User3"

	// # Clean up the list
	// RemoveOldestItem(ctx, key, colEmailHistory)
	// RemoveOldestItem(ctx, key, colPushHistory)
	// AppendHistoryItem(ctx, key, value, colEmailHistory)
	// AppendHistoryItem(ctx, key, value, colPushHistory)
	insertEventSuite(ctx, key, time.Now().Unix(), []byte(value))

	items, err := QueryEvent(ctx, key, "push", 10)
	if err != nil {
		log.Fatalf("Failed to query event: %+v", err)
	}
	log.Printf("items: %+v", len(items))

	items, err = QueryEvent(ctx, key, "email", 10)
	if err != nil {
		log.Fatalf("Failed to query event: %+v", err)
	}
	log.Printf("items: %+v", len(items))

	return
}

func insertEventSuite(ctx context.Context, key string, eventID int64, data []byte) (err error) {
	for i := 0; i < 3; i++ {
		var channels []string
		if i == 0 {
			channels = []string{"post-office", "push"}
		} else if i == 1 {
			channels = []string{"post-office", "email"}
		} else {
			channels = []string{"post-office", "push", "email"}
		}
		eventID := fmt.Sprintf("event-id-%d", eventID+int64(i))

		if err := insertEventWrapper(ctx, key, eventID, channels, data); err != nil {
			log.Fatalf("Failed to insert event: %+v", err)
		}
		if err := insertEventWrapper(ctx, key, eventID, channels, data); err != nil {
			log.Fatalf("Failed to insert event: %+v", err)
		}
	}
	return nil
}

func insertEventWrapper(ctx context.Context, key string, eventID string, stages []string, data []byte) (err error) {
	for _, stage := range stages {
		err = InsertEvent(ctx, key, eventID, stage, data)
		if err != nil {
			return
		}
	}
	return nil
}

func getUpdateExpression(col ColName) string {
	r := strings.NewReplacer("{col}", string(col))
	return r.Replace("SET #{col} = list_append(if_not_exists(#{col}, :emptyList), :historyItem)")
}

func AppendHistoryItem(ctx context.Context, key string, value string, col ColName) (err error) {
	// Update string
	attrMap := map[string]*string{}
	attrMap["#"+string(col)] = aws.String(string(col))

	resp, err := ddbClient.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Owner":   new(dynamodb.AttributeValue).SetS(key),
			"MsgType": new(dynamodb.AttributeValue).SetS(string(sortKeyV1Meta)),
		},
		UpdateExpression:         aws.String(getUpdateExpression(col)),
		ExpressionAttributeNames: attrMap,
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":historyItem": new(dynamodb.AttributeValue).SetL([]*dynamodb.AttributeValue{
				new(dynamodb.AttributeValue).SetS(value),
			}),
			":emptyList": new(dynamodb.AttributeValue).SetL([]*dynamodb.AttributeValue{}),
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		ReturnValues:           aws.String("ALL_NEW"),
	})
	if err != nil {
		log.Printf("Error: update item, %s", err)
		return err
	}
	fmt.Printf("result: %s\n", resp.GoString())
	return nil
}

func RemoveOldestItem(ctx context.Context, key string, col ColName) (err error) {
	resp, err := ddbClient.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Owner":   new(dynamodb.AttributeValue).SetS(key),
			"MsgType": new(dynamodb.AttributeValue).SetS(string(sortKeyV1Meta)),
		},
		UpdateExpression:    aws.String("REMOVE #History[1], #History[0]"),
		ConditionExpression: aws.String("size(#History) > :maxLen"),
		ExpressionAttributeNames: map[string]*string{
			"#History": aws.String(string(col)),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":maxLen": new(dynamodb.AttributeValue).SetN("5"),
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
	})
	if err != nil {
		log.Printf("Failed to truckate list: %+v", err)
	}
	fmt.Printf("trucate resp: %s\n", resp.GoString())
	return
}

func InitAws() *aws.Config {
	config := &aws.Config{
		Region: aws.String("us-west-2"),
	}
	return config
}
