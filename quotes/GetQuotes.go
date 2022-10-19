package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alikan97/lambda/models"
)

func GetQuotes() ([]models.AssetQuote, error) {
	symbolList := [5]string{"BTCUSDT", "IOTABTC", "ETHUSDT", "ADAUSDT", "XRPBTC"}

	result := make([]models.AssetQuote, 0)

	for _, s := range symbolList {
		request := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", s)
		resp, err := http.Get(request)

		if err != nil {
			fmt.Printf("Could not get data, %v", err)
			// Add logging
		}

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var marshaledObj models.AssetQuote
		errs := json.Unmarshal(responseData, &marshaledObj)
		if errs != nil {
			fmt.Printf("\n%v", errs)
		}

		result = append(result, marshaledObj)
	}

	fmt.Printf("%v", result)
	return result, nil
}
