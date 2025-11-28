package kafka

import (
	"sync"

	"github.com/segmentio/kafka-go"
)

type ConsumerManager struct {
	mu      sync.Mutex
	readers map[string]*kafka.Reader
	brokers []string
}

var cm = &ConsumerManager{
	readers: make(map[string]*kafka.Reader),
}

func InitConsumer(brokers []string) {
	cm.brokers = brokers
}

func GetReader(topic, groupId string) *kafka.Reader {
	key := groupId + "/" + topic
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if reader, ok := cm.readers[key]; ok {
		return reader
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cm.brokers,
		GroupID: groupId,
		Topic:   topic,
	})
	cm.readers[key] = reader
	return reader
}
