install:
	  go mod tidy

build:
	  go build src/main.go

run:
	  SQS_QUEUE_URL=https://sqs.us-east-1.amazonaws.com/XXXXXXXXXX/teste REGION=us-east-1 DYNAMODB_TABLE=your_table_name go run src/main.go