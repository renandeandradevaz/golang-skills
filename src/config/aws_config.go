package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func GetSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{Config: aws.Config{Region: aws.String(os.Getenv("REGION"))}, Profile: "default"}))
	return sess
}
