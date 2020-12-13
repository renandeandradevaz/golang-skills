install:
	  go mod tidy

build:
	  CGO_ENABLED=0 go build src/main.go

run:
	  SQS_QUEUE_URL=https://sqs.us-east-1.amazonaws.com/XXXXXXXXXX/teste REGION=us-east-1 DYNAMODB_TABLE=your_table_name go run src/main.go

build-image: build
	  docker build -t golang-skills .

run-image:
	  docker run -e AWS_ACCESS_KEY_ID=YOUR_AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=YOUR_AWS_SECRET_ACCESS_KEY -e AWS_DEFAULT_REGION=us-east-1 -e SQS_QUEUE_URL=https://sqs.us-east-1.amazonaws.com/XXXXXXXXXX/teste -e REGION=us-east-1 -e DYNAMODB_TABLE=your_table_name golang-skills