package main

import (
	"fmt"

	"github.com/alikan97/lambda.git/models"
	"github.com/alikan97/lambda.git/quotes"
	"github.com/google/uuid"
)

func main() {
	data, err := quotes.GetQuotes()
	if err != nil {
		fmt.Printf("%s", err)
	}

	MesageDto := models.MessageDTO[models.AssetQuote]{
		MessageType:    "ASSETQUOTE",
		MessageId:      uuid.New(),
		MessageContent: data,
	}

	fmt.Printf("MesageDto: %v\n", MesageDto)
}
