package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type tradecard struct {
	fromCurrency   string
	toCurrency     string
	finRate        float64
	alphaRate      float64
	boundaryHigher float64
	boundaryLower  float64
}

//

func newTradeCard(fromCurrency string, toCurrency string, finRate float64, alphaRate float64, boundaryHigher float64, boundaryLower float64) *tradecard {
	t := tradecard{fromCurrency: fromCurrency}
	t = tradecard{toCurrency: toCurrency}
	t = tradecard{finRate: finRate}
	t = tradecard{alphaRate: alphaRate}
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
	ftokenPtr := flag.String("finntoken", "", "Finnhub token")
	atokenPtr := flag.String("alphatoken", "", "AlphaVantage token")

	finnhubToken := *ftokenPtr
	alphavantageToken := *atokenPtr

	fRate := queryFinnhub(*fromPtr, *toPtr, finnhubToken)
	aRate := queryAlphaVantage(*fromPtr, *toPtr, alphavantageToken)

	//Build Trade Card
	currentTrade := tradecard{fromCurrency: *fromPtr, toCurrency: *toPtr, finRate: fRate, alphaRate: aRate, boundaryHigher: *hbPtr, boundaryLower: *lbPtr}
	fmt.Println("Trade Card Created")
	fmt.Println(currentTrade.fromCurrency)
	fmt.Println(currentTrade.toCurrency)

	//Scanner - card, frequency, tokens, repetitions
	marketScan(currentTrade, finnhubToken, alphavantageToken, 60, 480)

	//Alert - End of Program Alert
	alert("end")

}

func queryAlphaVantage(fromcurrency string, tocurrency string, token string) float64 {
	fmt.Println("Query Alpha")
	var num float64
	num = 1.35672
	return num
}

func queryFinnhub(fromcurrency string, tocurrency string, token string) float64 {
	fmt.Println("Query Finnhub")

	//Service String
	service := "https://finnhub.io/api/v1/forex/rates?base=" + fromcurrency + "&token=" + token

	client := http.Client{}
	request, err := http.NewRequest("GET", service, nil)
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
	var num float64
	num = 1.35672
	return num
}

func marketScan(currentTrade tradecard, finnhubToken string, alphavantageToken string, frequency int, repetition int) {

	//loop at frequency as defined in seconds and repeat for repitition

	for i := 0; i < repetition; i++ {

		//query prices and check card
		fRate := queryFinnhub(currentTrade.fromCurrency, currentTrade.toCurrency, finnhubToken)
		aRate := queryAlphaVantage(currentTrade.fromCurrency, currentTrade.toCurrency, alphavantageToken)

		if fRate > currentTrade.boundaryHigher {
			fmt.Println("Finn Above")
			alert("trade")
		} else if fRate < currentTrade.boundaryLower {
			fmt.Println("Finn below")
			alert("trade")
		} else {
			fmt.Println("Nothing")
		}

		if aRate > currentTrade.boundaryHigher {
			fmt.Println("Alpha Above")
			alert("trade")

		} else if aRate < currentTrade.boundaryLower {
			fmt.Println("Alpha below")
			alert("trade")

		} else {
			fmt.Println("Nothing")
		}

		//sleep frequency
		time.Sleep(time.Duration(frequency) * time.Second)

	}

}

func alert(alertType string) {

	if strings.Compare(alertType, "end") == 0 {
		fmt.Println("Progran Ending")

	} else {
		fmt.Println("Trade Possibly")
	}

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
