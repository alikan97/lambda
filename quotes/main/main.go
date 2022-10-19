package main

import (
	"fmt"

	"github.com/alikan97/lambda/models"
	"github.com/alikan97/lambda/quotes"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

func HandleRequest() (models.MessageDTO[models.AssetQuote], error) {
	data, err := quotes.GetQuotes()
	if err != nil {
		fmt.Printf("%s", err)
		return models.MessageDTO[models.AssetQuote]{}, err
	}

	MesageDto := models.MessageDTO[models.AssetQuote]{
		MessageType:    "ASSETQUOTE",
		MessageId:      uuid.New(),
		MessageContent: data,
	}

	fmt.Printf("Successfully retrieved asset quotes data from Binance")
	return MesageDto, nil
}

func main() {
	lambda.Start(HandleRequest)
}
