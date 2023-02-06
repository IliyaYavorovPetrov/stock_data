package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const apiName = "https://twelvedata.com"
const noAnswerQuote = "NO_ANSWER_QUOTE"
const noAnswerPrice = -1

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

	if quote.QuoteString == "" {
		return noAnswerQuote
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
		return noAnswerPrice
	}

	priceFloat64, err := strconv.ParseFloat(priceString.PriceString, 64)
	if err != nil {
		fmt.Println(err)
	}

	return priceFloat64
}
