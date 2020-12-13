package worker

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/renandeandradevaz/golang-skills/src/config"
	"github.com/renandeandradevaz/golang-skills/src/constants"
	"github.com/renandeandradevaz/golang-skills/src/service"
)

func InitWorker() {
	sess := config.GetSession()
	service.InitSqsClient(sess)

	for i := 1; i <= constants.NumberOfWorkers; i++ {
		go asyncJob()
	}
}

func asyncJob() {
	chnMessages := make(chan *sqs.Message, constants.MaxNumberOfMessages)
	go service.PollMessages(chnMessages)
	for message := range chnMessages {
		handleMessage(message)
	}
}

func handleMessage(msg *sqs.Message) {

	fmt.Println(*msg.Body)
	service.DeleteMessage(msg)
}
