package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
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
		return 0
	}

	priceFloat64, err := strconv.ParseFloat(priceString.PriceString, 64)
	if err != nil {
		fmt.Println(err)
	}

	return priceFloat64
}

func main() {
	tickers := [7]string{"AAPL", "AMZN", "TSLA", "META", "PFE", "KO", "WMT"}
	apiKey := "e808bc63e1de4120a2690e7d4a447156"

	c := cron.New()
	_, err := c.AddFunc("@every 20m", func() {
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
	})
	if err != nil {
		fmt.Println(err)
	}

	c.Start()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Type \"q\" to quit..")
		comm, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		if strings.Compare(comm, "q") == 1 {
			c.Stop()
			break
		}
	}
}
