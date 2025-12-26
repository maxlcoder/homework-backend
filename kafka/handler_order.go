package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type OrderCreatedMsg struct {
	OrderID string `json:"orderId"`
	UserID  string `json:"userId"`
}

func OrderCreatedHandler(msg kafka.Message) error {
	var m OrderCreatedMsg
	if err := json.Unmarshal(msg.Value, &m); err != nil {
		return err
	}
	fmt.Println("订单创建事件 ->", m.OrderID, m.UserID)

	// TODO: 调用 service 层处理业务逻辑

	return nil
}
