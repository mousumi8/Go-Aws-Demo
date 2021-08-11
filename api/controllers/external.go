package controllers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)
var AccessKeyID string
var SecretAccessKey string
var MyRegion string
func ConnectAws(accessKeyID,secretAccessKey string) *session.Session {
	AccessKeyID = accessKeyID
	SecretAccessKey = secretAccessKey
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
func (s *Server) LoadAwsRegion( awsRegion string){
	MyRegion = awsRegion
}

