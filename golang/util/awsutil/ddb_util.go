package awsutil

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DdbVal() *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{}
}
