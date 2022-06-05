package handler

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var awsSessionStore atomic.Value
var awsSqsClientStore atomic.Value
var awsDdbClientStore atomic.Value

func getSQSClient() *sqs.SQS {
	svc, ok := awsSqsClientStore.Load().(*sqs.SQS)
	if ok {
		return svc
	}
	svc = newSQSClient()
	awsSqsClientStore.Store(svc)
	return svc
}

func getDdbClient() *dynamodb.DynamoDB {
	svc, ok := awsDdbClientStore.Load().(*dynamodb.DynamoDB)
	if ok {
		return svc
	}
	svc = newDdbClient()
	awsDdbClientStore.Store(svc)
	return svc
}

func newSQSClient() *sqs.SQS {
	return sqs.New(getAwsSession())
}

func newDdbClient() *dynamodb.DynamoDB {
	return dynamodb.New(getAwsSession())
}

func val() *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{}
}

func getAwsSession() *session.Session {
	sess, ok := awsSessionStore.Load().(*session.Session)
	if ok {
		// Check session is valid
		t, err := sess.Config.Credentials.ExpiresAt()
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "ProviderNotExpirer" {
				// it's ok like this.
				t = time.Now().Add(time.Hour)
				err = nil
			}
		}
		if err != nil {
			log.Printf("failed to get session expiration, ignoring...\n err: %+v\n", err)
			ok = false
		} else {
			if t.Before(time.Now().Add(time.Minute)) {
				log.Println("session expiring soon or expired, renewing...")
				ok = false
			}
		}
	}
	if !ok || sess == nil {
		if sess == nil {
			if isDebugMode {
				log.Println("session is nil, creating new one...")
			}
		}
		var err error
		sess, err = newAwsSession()
		if err != nil {
			log.Fatalf("unable to initialize aws session, err: %+v\n", err)
		}
		awsSessionStore.Store(sess)
	}

	return sess
}

func newAwsSession() (*session.Session, error) {
	awsConfig := aws.NewConfig().WithRegion("us-west-2")
	return session.NewSession(awsConfig)
}
