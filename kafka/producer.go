package kafka

import (
	"context"
	"sync"

	"github.com/segmentio/kafka-go"
)

type ProducerManager struct {
	mu      sync.Mutex
	writers map[string]*kafka.Writer
	brokers []string
}

var pm = &ProducerManager{
	writers: make(map[string]*kafka.Writer),
}

func InitProducer(brokers []string) {
	pm.brokers = brokers
}

// GetWriter 懒加载：按 topic 创建 Writer (防止重复 new)
func GetWriter(topic string) *kafka.Writer {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if writer, ok := pm.writers[topic]; ok {
		return writer
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP(pm.brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	pm.writers[topic] = writer
	return writer
}

// SendSync 同步发送
func SendSync(topic string, value []byte) error {
	writer := GetWriter(topic)
	return writer.WriteMessages(context.Background(), kafka.Message{
		Value: value,
	})
}

// SendAsync 异步发送
func SendAsync(topic string, value []byte) {
	writer := GetWriter(topic)
	go writer.WriteMessages(context.Background(), kafka.Message{
		Value: value,
	})
}
