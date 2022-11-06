package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/alikan97/lambda/models"
	pb "github.com/alikan97/lambda/proto"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/grpc"
)

func HandleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	serverAddress := flag.String("address", "172.30.2.10:8085", "the server address")

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	c := pb.NewCryptoClient(conn)

	for _, msg := range sqsEvent.Records {
		message := &models.MessageDTO[any]{}

		err := json.Unmarshal([]byte(msg.Body), message)
		if err != nil {
			fmt.Errorf("Failed to unmarshal: %v", err)
			return err
		}

		if message.MessageType == "ASSETQUOTE" {
			for _, v := range message.MessageContent {
				jsonBody, err := json.Marshal(v)
				if err != nil {
					fmt.Printf("Error unmarshaling json")
				}

				AQ := models.AssetQuote{}

				if err := json.Unmarshal(jsonBody, &AQ); err != nil {
					fmt.Printf("Error unmarshaling to models.AssetQuote")
				}

				newPrice, err := strconv.ParseFloat(AQ.Price, 64)
				if err != nil {
					fmt.Errorf("Error trying to parse string to float: %v", err)
				}
				fmt.Printf("SymbolName: %s, Price: %v", AQ.SymbolName, newPrice)
				_, rpcErr := c.UpdateQuotes(ctx, &pb.UpdateQuoteReq{Symbol: AQ.SymbolName, UpdatedPrice: newPrice})

				if rpcErr != nil {
					fmt.Errorf("Error occurred trying to send update for message id: %s. Reason: %v", message.MessageId, rpcErr)
					return rpcErr
				}
			}
			return nil
		}

		if message.MessageType == "RECENTTRADES" {
			msgArr := make([]*pb.AddRecentTradeReq, 0)

			for _, v := range message.MessageContent {
				jsonBody, err := json.Marshal(v)
				if err != nil {
					fmt.Errorf("Error unmarshaling json")
				}

				RT := models.RecentTradesDTO{}

				if err := json.Unmarshal(jsonBody, &RT); err != nil {
					fmt.Printf("Error unmarshalling to models.RecentTradesDto")
				}

				formattedMsg := &pb.AddRecentTradeReq{
					AssetName: RT.AssetName,
					AssetCode: RT.AssetCode,
					Price:     RT.Price,
					Quantity:  RT.Quantity,
					Time:      RT.Time,
				}

				msgArr = append(msgArr, formattedMsg)
			}

			fmt.Printf("Value: %v", msgArr)
			_, rpcErr := c.AddRecentTrade(ctx, &pb.AddRecentTradeReqMulti{RecentTrade: msgArr})

			if rpcErr != nil {
				fmt.Errorf("Error adding recent trade for message id: %s. Reason: %v", message.MessageId, rpcErr)
				return rpcErr
			}
			return nil
		}
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
