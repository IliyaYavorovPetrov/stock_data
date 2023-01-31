package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func getStockQuote(ticker string, apiKey string) string {
	url := "https://api.twelvedata.com/quote?symbol=" + ticker + "&apikey=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	type QuoteStruct struct {
		QuoteString string `json:"name"`
	}

	var quote QuoteStruct
	if err := json.Unmarshal(data, &quote); err != nil {
		fmt.Println(err)
	}

	return quote.QuoteString
}

func getStockPrice(ticker string, apiKey string) float64 {
	url := "https://api.twelvedata.com/price?symbol=" + ticker + "&apikey=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	type PriceStruct struct {
		PriceString string `json:"price"`
	}

	var priceString PriceStruct
	if err := json.Unmarshal(data, &priceString); err != nil {
		fmt.Println(err)
	}

	if priceString.PriceString == "" {
		fmt.Println("Rate limit is reached, please wait")
		return 0
	}

	priceFloat64, err := strconv.ParseFloat(priceString.PriceString, 64)
	if err != nil {
		fmt.Println(err)
	}

	return priceFloat64
}

func main() {
	tickers := [8]string{"AAPL", "JNJ", "AMZN", "TSLA", "META", "PFE", "KO", "WMT"}
	apiKey := "e808bc63e1de4120a2690e7d4a447156"

	timeTicker := time.NewTicker(90 * time.Second)
	timeQuit := make(chan struct{})

	go func() {
		for {
			select {
			case <-timeTicker.C:
				var wg sync.WaitGroup
				wg.Add(len(tickers))

				for i, x := range tickers {
					ticker := x
					go func(i int) {
						defer wg.Done()
						quote := getStockQuote(ticker, apiKey)
						price := getStockPrice(ticker, apiKey)
						fmt.Printf("%s: %f\n", quote, price)
					}(i)
				}

				wg.Wait()
			case <-timeQuit:
				timeTicker.Stop()
			}
		}
	}()
}
