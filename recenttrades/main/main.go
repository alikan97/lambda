package main

import (
	"fmt"

	"github.com/alikan97/lambda.git/models"
	"github.com/alikan97/lambda.git/recenttrades"
	"github.com/google/uuid"
)

func main() {
	data, err := recenttrades.GetRecentTrades()
	if err != nil {
		fmt.Printf("%s", err)
	}

	MesageDto := models.MessageDTO[models.RecentTradesDTO]{
		MessageType:    "ASSETQUOTE",
		MessageId:      uuid.New(),
		MessageContent: data,
	}

	fmt.Printf("MesageDto: %v\n", MesageDto)
}
