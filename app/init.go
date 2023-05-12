package app

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/esimov/xm/app/models"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB       *gorm.DB
	Route    *gin.Engine
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
}

func (s *Server) Init(c *config.Config) error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.UserName, c.Password, c.HostName, c.Port, c.DB)
	fmt.Println(dsn)
	s.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	err = models.Load(s.DB)
	if err != nil {
		log.Fatalf("Server error, %v", err)
	}

	err = s.InitRoutes(c)
	if err != nil {
		log.Fatalf("Kafka error, %v", err)
	}

	return nil
}

func (s *Server) Send(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := s.Producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func (s *Server) Receive(topic string) {
	partitionList, err := s.Consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, _ := s.Consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				fmt.Println(string(message.Value))
			}
		}(pc)
	}
}
