package worker

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/guregu/dynamo"
	"github.com/renandeandradevaz/golang-skills/src/client"
	"github.com/renandeandradevaz/golang-skills/src/config"
	"github.com/renandeandradevaz/golang-skills/src/constants"
	"github.com/renandeandradevaz/golang-skills/src/structs"
)

var (
	dynamoClient dynamo.Table
)

func InitWorker() {
	sess := config.GetSession()
	client.InitSqsClient(sess)
	dynamoClient = client.InitDynamoClient(sess)

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

	key := messageFromSqs + "-" + ip.Ip

	var myDynamoModel structs.MyDynamoModel
	err = dynamoClient.Get("key", key).One(&myDynamoModel)
	if err == nil {
		myDynamoModel.Count++
	} else {
		myDynamoModel = structs.MyDynamoModel{}
		myDynamoModel.Key = key
		myDynamoModel.Msg = "Message for key: " + key
		myDynamoModel.Count = 1
	}

	err = dynamoClient.Put(myDynamoModel).Run()
	if err != nil {
		fmt.Printf("Unable to put object on dynamodb %v\n", err)
		return
	}

	fmt.Printf("Saving object on dynamodb %v\n", myDynamoModel)

	client.DeleteSqsMessage(msg)
}
