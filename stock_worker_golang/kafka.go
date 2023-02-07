package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func setUpKafka() (*kafka.Producer, *kafka.Consumer) {
	prod, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	cons, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "default",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	return prod, cons
}
