package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

var c ConfigProvider
var region string
var profile string

// ConfigProvider ...
type ConfigProvider struct{}

// ClientConfig ...
func (cp ConfigProvider) ClientConfig(serviceName string, cfgs ...*aws.Config) client.Config {
	return client.Config{
		Config: &aws.Config{
			Region:     aws.String("us-east-1"),
			MaxRetries: aws.Int(0),
		},
	}
}

// https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.UpdateItem
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	log.Println("Creating session in Handler()")
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	Profile: profile,
	// 	Config: aws.Config{
	// 		Region: aws.String(region),
	// 	},
	// }))
	// db := dynamodb.New(sess)
	// op, err := db.PutItem()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
