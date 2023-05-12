package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

func Producer(brokerList []string) sarama.SyncProducer {
	// For the data collector, we are looking for strong consistency semantics.
	// Because we don't change the flush settings, sarama will try to produce messages
	// as fast as possible to keep latency low.
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}

func Consumer(brokerList []string) sarama.Consumer {
	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		log.Fatalln("Failed to create Sarama consumer:", err)
	}
	return consumer
}
