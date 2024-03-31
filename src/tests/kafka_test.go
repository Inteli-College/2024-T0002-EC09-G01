package testing

import (
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func TestKafkaConsumer_Consume(t *testing.T) {
	configMap := &kafka.ConfigMap{}
	topics := []string{"test_topic"}

	msgChan := make(chan *kafka.Message)

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		t.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		t.Fatalf("Error subscribing to topics: %v", err)
	}

	go func() {
		for {
			select {
			case <-time.After(100 * time.Millisecond):
				return
			case ev := <-consumer.Events():
				switch e := ev.(type) {
				case kafka.AssignedPartitions:
					consumer.Assign(e.Partitions)
				case kafka.RevokedPartitions:
					consumer.Unassign()
				case *kafka.Message:
					msgChan <- e
				case kafka.Error:
					t.Errorf("Kafka error: %v", e)
				}
			}
		}
	}()

	for i := 0; i < 3; i++ {
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topics[0], Partition: 0},
			Value:          []byte("Test message"),
		}
		msgChan <- message
	}

	close(msgChan)
}