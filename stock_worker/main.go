package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/robfig/cron/v3"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

const apiName = "https://twelvedata.com"
const noAnswerQuote = "NO_ANSWER_QUOTE"
const noAnswerPrice = -1

func main() {
	tickers := [4]string{"AAPL", "AMZN", "META", "TSLA"}
	apiKey := "e808bc63e1de4120a2690e7d4a447156"
	topic := "stock"

	prod, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer prod.Close()

	go func() {
		for e := range prod.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	prod.Flush(15 * 1000)

	c := cron.New()
	_, err = c.AddFunc("0,51 * * * *", func() {
		var wg sync.WaitGroup
		wg.Add(len(tickers))

		for i, x := range tickers {
			ticker := x
			go func(i int) {
				defer wg.Done()
				currTime := time.Now()
				quote := getStockQuote(ticker, apiKey)
				price := getStockPrice(ticker, apiKey)

				if quote != noAnswerQuote && price != noAnswerPrice {
					fmt.Printf("[%s] %s: %.2f\n", currTime.Format(time.RFC1123), quote, price)
					msg := "Hello from Go"

					err := prod.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(msg),
					}, nil)
					if err != nil {
						fmt.Println(err)
					}

				} else {
					fmt.Printf("No answer received from %s for %s\n", apiName, ticker)
				}
			}(i)
		}

		wg.Wait()
		fmt.Println("Type \"q\" to quit...")
	})
	if err != nil {
		fmt.Println(err)
	}

	c.Start()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Type \"q\" to quit...")
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
