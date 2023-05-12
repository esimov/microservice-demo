package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func Send(producer sarama.SyncProducer, topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
