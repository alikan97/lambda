package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alikan97/lambda/models"
	"github.com/alikan97/lambda/recenttrades"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

func HandleRequest() error {
	data, err := recenttrades.GetRecentTrades()

	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	qurl := os.Getenv("qurl")

	MesageDto, _ := json.Marshal(models.MessageDTO[models.RecentTradesDTO]{
		MessageType:    "RECENTTRADES",
		MessageId:      uuid.New(),
		MessageContent: data,
	})

	fmt.Printf("Successfully retrieved recent trades data from Binance")

	err = SendMessage(sess, qurl, string(MesageDto))

	if err != nil {
		fmt.Printf("Got an error while trying to send message to queue: %v", err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}

func SendMessage(sess *session.Session, queueUrl string, messageBody string) error {
	sqsClient := sqs.New(sess)

	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: aws.String(messageBody),
	})

	return err
}
