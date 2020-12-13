package worker

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/renandeandradevaz/golang-skills/src/client"
	"github.com/renandeandradevaz/golang-skills/src/config"
	"github.com/renandeandradevaz/golang-skills/src/constants"
	"github.com/renandeandradevaz/golang-skills/src/structs"
)

func InitWorker() {
	sess := config.GetSession()
	client.InitSqsClient(sess)

	for i := 1; i <= constants.NumberOfWorkers; i++ {
		go asyncJob()
	}
}

func asyncJob() {
	chnMessages := make(chan *sqs.Message, constants.MaxNumberOfMessages)
	go client.PollMessagesFromSqs(chnMessages)
	for message := range chnMessages {
		handleMessage(message)
	}
}

func handleMessage(msg *sqs.Message) {

	messageFromSqs := *msg.Body
	fmt.Println(messageFromSqs)

	responseJson, err := client.GetRequest("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Printf("Unable to make http request %v\n", err)
		return
	}

	ip := structs.Ip{}
	err = json.Unmarshal(responseJson, &ip)
	if err != nil {
		fmt.Printf("Unable to Unmarshal json %v\n", err)
		return
	}

	fmt.Println(ip.Ip)

	

	client.DeleteSqsMessage(msg)
}
