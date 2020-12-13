package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

func InitDynamoClient(sess *session.Session) dynamo.Table {
	return dynamo.New(sess, &aws.Config{Region: aws.String(os.Getenv("REGION"))}).Table(os.Getenv("DYNAMODB_TABLE"))
}
