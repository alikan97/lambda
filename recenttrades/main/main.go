package main

import (
	"fmt"

	"github.com/alikan97/lambda/models"
	"github.com/alikan97/lambda/recenttrades"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

func HandleRequest() (models.MessageDTO[models.RecentTradesDTO], error) {
	data, err := recenttrades.GetRecentTrades()
	if err != nil {
		fmt.Printf("%s", err)
		return models.MessageDTO[models.RecentTradesDTO]{}, err
	}

	MesageDto := models.MessageDTO[models.RecentTradesDTO]{
		MessageType:    "RECENTTRADES",
		MessageId:      uuid.New(),
		MessageContent: data,
	}

	fmt.Printf("Successfully retrieved recent trades data from Binance")
	return MesageDto, nil
}

func main() {
	lambda.Start(HandleRequest)
}
