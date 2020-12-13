package structs

type Ip struct {
	Ip string `json:"ip"`
}

type MyDynamoModel struct {
	Key   string `dynamo:"key"`
	Msg   string `dynamo:"message"`
	Count int    `dynamo:"count"`
}
