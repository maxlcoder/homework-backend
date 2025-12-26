package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/segmentio/kafka-go"
)

type HandleFunc func(msg kafka.Message) error

type ConsumerManager struct {
	mu       sync.Mutex
	readers  map[string]*kafka.Reader
	handlers map[string]HandleFunc
	brokers  []string

	wg     sync.WaitGroup
	stopCh chan struct{}
}

var cm = &ConsumerManager{
	readers:  make(map[string]*kafka.Reader),
	handlers: make(map[string]HandleFunc),
	stopCh:   make(chan struct{}),
}

func InitConsumer(brokers []string) {
	cm.brokers = brokers
}

func (cm *ConsumerManager) getReader(topic, groupId string) *kafka.Reader {
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

func RegisterAllConsumers() {
	RegisterConsumer("order-group", "order-created", OrderCreatedHandler)
}

// 注册 topic + group 对应的 handler
func RegisterConsumer(groupId, topic string, handler HandleFunc) {
	key := groupId + "/" + topic
	cm.handlers[key] = handler
}

// 有多少 handler 就启动多少
func StartConsumers() {
	for key, handler := range cm.handlers {
		parts := strings.Split(key, "/")
		groupId, topic := parts[0], parts[1]
		reader := cm.getReader(topic, groupId)
		cm.wg.Add(1)
		go cm.runConsumer(reader, handler)
	}
	go cm.waitForExit()
}

// 执行消费
func (cm *ConsumerManager) runConsumer(reader *kafka.Reader, handler HandleFunc) {
	defer cm.wg.Done()
	for {
		select {
		case <-cm.stopCh:
			_ = reader.Close()
			return
		default:
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("reader error:", err)
				continue
			}
			// 并发执行 handler
			go func(m kafka.Message) {
				if err := handler(m); err != nil {
					log.Println("handler error:", err)
				}
			}(msg)
		}
	}
}

func (cm *ConsumerManager) waitForExit() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("Kafka shutting down...")

	close(cm.stopCh)
	cm.wg.Wait()
	log.Println("Kafka shutdown complete")
}
