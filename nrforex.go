package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type tradecard struct {
	fromCurrency   string
	toCurrency     string
	rate           float64
	boundaryHigher float64
	boundaryLower  float64
}

//

func newTradeCard(fromCurrency string, toCurrency string, rate float64, boundaryHigher float64, boundaryLower float64) *tradecard {
	t := tradecard{fromCurrency: fromCurrency}
	t = tradecard{toCurrency: toCurrency}
	t = tradecard{rate: rate}
	t = tradecard{boundaryHigher: boundaryHigher}
	t = tradecard{boundaryLower: boundaryLower}
	return &t
}

//

func main() {

	fromPtr := flag.String("from", "GBP", "From Currency")
	toPtr := flag.String("To", "USD", "To Currency")
	hbPtr := flag.Float64("high", 1.000, "Higher Boundary")
	lbPtr := flag.Float64("low", 1.000, "Lower Boundary")

	client := http.Client{}
	request, err := http.NewRequest("GET", "https://finnhub.io/api/v1/forex/rates?base=USD&token=", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

	fmt.Println("HERE")

	currentTrade := tradecard{fromCurrency: *fromPtr, toCurrency: *toPtr, rate: 1.32456, boundaryHigher: *hbPtr, boundaryLower: *lbPtr}
	fmt.Println(currentTrade.fromCurrency)

}

func queryAlphaVantage(fromcurrency string, tocurrency string) float32 {
	fmt.Println("Query Alpha")
	var num float32
	num = 1.35672
	return num
}

func queryFinnhub(fromcurrency string, tocurrency string) float32 {
	fmt.Println("Query Finnhub")
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://finnhub.io/api/v1/forex/rates?base=USD&token=", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

	//test text
	var num float32
	num = 1.35672
	return num
}

//Take input of currency pair and boundary values and time range for boudaries
//Scan market every minute checking for breach of the boundary
//On boundary breach set off alert

//Functions
//1 Query Finnhub
//2 Query Alphavantage
//3 unpack Finnhub response
//4 unpack alphavantage reponse
//5 Check for alert breach
