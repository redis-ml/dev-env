package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/redisliu/dev-env/golang/cmd/sqsfanoutworker/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
