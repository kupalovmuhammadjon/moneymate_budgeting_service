package kafka

import (
	pb "budgeting_service/genproto/budgeting_service"
	"budgeting_service/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
)

type IKafka interface {
	Close()
	ProduceMessage(string, string) error
	ConsumeMessages([]string, services.IServiceManager) error
}

type kafka struct {
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup
}

type consumerGroupHandler struct {
	iServiceManager services.IServiceManager
}

func (h *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Received message: %s\n", string(message.Value))
		switch message.Topic {
		case "transaction_created":
			req := pb.CreateTransaction{}
			err := json.Unmarshal(message.Value, &req)
			if err != nil {
				log.Println("failed to unmarshal transaction message", err)
				continue
			}
			_, err = h.iServiceManager.TransactionService().Create(context.Background(), &req)
			if err != nil {
				log.Println("failed to create transaction", err)
			}
		case "budget_updated":
			req := pb.Budget{}
			err := json.Unmarshal(message.Value, &req)
			if err != nil {
				log.Println("failed to unmarshal budget message", err)
				continue
			}
			_, err = h.iServiceManager.BudgetService().Update(context.Background(), &req)
			if err != nil {
				log.Println("failed to update budget", err)
			}
		case "goal_progress_updated":
			req := pb.Goal{}
			err := json.Unmarshal(message.Value, &req)
			if err != nil {
				log.Println("failed to unmarshal goal message", err)
				continue
			}
			_, err = h.iServiceManager.GoalService().Update(context.Background(), &req)
			if err != nil {
				log.Println("failed to update goal", err)
			}
		case "notification_created":
		}
		session.MarkMessage(message, "")
	}

	return nil
}

func NewIKafka() (IKafka, error) {
	producer, err := newKafkaProducer()
	if err != nil {
		return nil, err
	}

	consumer, err := newKafkaConsumer()
	if err != nil {
		return nil, err
	}

	return &kafka{
		producer: producer,
		consumer: consumer,
	}, nil
}

func (k *kafka) ProduceMessage(topic, value string) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}
	_, _, err := k.producer.SendMessage(message)
	if err != nil {
		return err
	}

	return nil
}

func (k *kafka) ConsumeMessages(topics []string, iServiceManager services.IServiceManager) error {
	handler := consumerGroupHandler{
		iServiceManager: iServiceManager,
	}

	for {
		err := k.consumer.Consume(context.Background(), topics, &handler)
		if err != nil {
			return err
		}
	}
}

func (k *kafka) Close() {
	k.producer.Close()
	k.consumer.Close()
}

func newKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Producer.Return.Successes = true
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	producer, err := sarama.NewSyncProducer([]string{":29092"}, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func newKafkaConsumer() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Producer.Return.Successes = true
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	consumer, err := sarama.NewConsumerGroup([]string{":29092"}, "budgeting-service", config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
