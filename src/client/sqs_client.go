package client

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/renandeandradevaz/golang-skills/src/constants"
	"os"
	"time"
)

var (
	sqsClient *sqs.SQS
)

func InitSqsClient(sess *session.Session) {
	sqsClient = sqs.New(sess)
}

func PollMessagesFromSqs(chn chan<- *sqs.Message) {

	for {
		output, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(os.Getenv("SQS_QUEUE_URL")),
			MaxNumberOfMessages: aws.Int64(constants.MaxNumberOfMessages),
			WaitTimeSeconds:     aws.Int64(20),
		})

		if err != nil {
			fmt.Printf("failed to fetch sqs message %v\n", err)
			time.Sleep(10 * time.Second)
		} else {
			for _, message := range output.Messages {
				chn <- message
			}
		}
	}
}

func DeleteSqsMessage(msg *sqs.Message) {

	sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("SQS_QUEUE_URL")),
		ReceiptHandle: msg.ReceiptHandle,
	})
}
