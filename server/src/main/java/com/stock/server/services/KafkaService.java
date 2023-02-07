package com.stock.server.services;

import org.springframework.kafka.annotation.KafkaListener;

public class KafkaService {
    @KafkaListener(topics = "stock", groupId = "default")
    void listenForEvents(String jsonData) {
        System.out.println("Kafka listener received: " + jsonData);
    }
}
