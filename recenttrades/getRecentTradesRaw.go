package recenttrades

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/alikan97/lambda.git/models"
)

func GetRecentTrades() ([]models.RecentTradesDTO, error) {
	symbolList := [5]string{"BTCUSDT", "IOTABTC", "ETHUSDT", "ADAUSDT", "XRPBTC"}

	result := make([]models.RecentTradesDTO, 0)

	for _, s := range symbolList {
		request := fmt.Sprintf("https://api.binance.com/api/v3/trades?symbol=%s&limit=10", s)
		resp, err := http.Get(request)

		if err != nil {
			fmt.Printf("Could not get data, %v", err)
			// Add logging
		}

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var marshaledObj []models.RawRecentTrades
		errs := json.Unmarshal(responseData, &marshaledObj)
		if errs != nil {
			fmt.Printf("\n%v", errs)
		}

		for _, val := range marshaledObj {
			Price, _ := strconv.ParseInt(val.Price, 0, 32)
			Quantity, _ := strconv.ParseFloat(val.Quantity, 64)

			result = append(result, models.RecentTradesDTO{
				AssetName: s,
				AssetCode: "CRYP",
				Price:     uint32(Price),
				Quantity:  Quantity,
				Time:      uint64(val.Time),
			})
		}
	}

	return result, nil
}
