package main

import (
	"bufio"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"strings"
)

func main() {
	tickers := [4]string{"AAPL", "AMZN", "META", "TSLA"}
	apiKey := "e808bc63e1de4120a2690e7d4a447156"
	topic := "stock"

	l := log.New(os.Stdout, "stock_worker_golang ", log.LstdFlags)

	// Set up producer Kafka
	prod, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer prod.Close()

	prod.Flush(15 * 1000)

	// Create the cron job
	type StockDataStruct struct {
		StockDataQuote string  `json:"quote"`
		StockDataPrice float64 `json:"price"`
	}

	c := cron.New()
	_, err = c.AddFunc("0,30 * * * *", func() {
		for _, t := range tickers {
			ticker := t
			go func() {
				quote := getStockQuote(ticker, apiKey)
				price := getStockPrice(ticker, apiKey)

				if quote != noAnswerQuote && price != noAnswerPrice {
					l.Printf("%s: %.2f\n", quote, price)

					data := StockDataStruct{
						StockDataQuote: quote,
						StockDataPrice: price,
					}

					dataJson, _ := json.Marshal(data)

					err := prod.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
						Value:          []byte(string(dataJson)),
					}, nil)
					if err != nil {
						l.Println(err)
					}

				} else {
					l.Printf("No answer received from %s for %s\n", apiName, ticker)
				}
			}()
		}
	})
	if err != nil {
		l.Println(err)
	}

	// Start the stock worker
	c.Start()
	reader := bufio.NewReader(os.Stdin)
	for {
		l.Println("Type \"q\" to quit...")
		comm, err := reader.ReadString('\n')
		if err != nil {
			l.Println(err)
		}

		if strings.Compare(comm, "q") == 1 {
			c.Stop()
			break
		}
	}
}
