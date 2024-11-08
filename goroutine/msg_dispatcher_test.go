package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// Message 表示要处理的消息
type Message struct {
	ID      int
	Content string
}

// Handler 表示消息处理程序的接口
type Handler interface {
	HandleMessage(Message)
}

// ConditionHandler 表示带有条件的消息处理程序
type ConditionHandler struct {
	Condition func(Message) bool
	Handler   Handler
}

// Dispatcher 表示消息调度器
type Dispatcher struct {
	mu       sync.Mutex
	handlers []ConditionHandler
}

// NewDispatcher 创建一个新的消息调度器
func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

// AddHandler 添加一个带有条件的消息处理程序
func (d *Dispatcher) AddHandler(condition func(Message) bool, handler Handler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers = append(d.handlers, ConditionHandler{Condition: condition, Handler: handler})
}

// Dispatch 根据消息的条件选择合适的处理程序并处理消息
func (d *Dispatcher) Dispatch(message Message) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, conditionHandler := range d.handlers {
		if conditionHandler.Condition(message) {
			conditionHandler.Handler.HandleMessage(message)
			return
		}
	}

	// 如果没有匹配的处理程序，默认处理方式
	fmt.Printf("No matching handler found for Message %d: %s\n", message.ID, message.Content)
}

// SimpleHandler 是一个简单的消息处理程序
type SimpleHandler struct {
	Name string
}

// HandleMessage 实现了 Handler 接口
func (h *SimpleHandler) HandleMessage(message Message) {
	fmt.Printf("[%s] Handling Message %d: %s\n", h.Name, message.ID, message.Content)
}

func TestMsgDispatcher(t *testing.T) {
	// 创建消息调度器
	dispatcher := NewDispatcher()

	// 添加带有条件的消息处理程序
	dispatcher.AddHandler(func(message Message) bool {
		// 条件：消息的 ID 是偶数
		return message.ID%2 == 0
	}, &SimpleHandler{Name: "EvenHandler"})

	dispatcher.AddHandler(func(message Message) bool {
		// 条件：消息的内容包含 "important"
		return message.Content == "important"
	}, &SimpleHandler{Name: "ImportantHandler"})

	// 模拟产生一些消息并进行分发
	messages := []Message{
		{ID: 1, Content: "Hello"},
		{ID: 2, Content: "World"},
		{ID: 3, Content: "This is important"},
	}

	for _, message := range messages {
		dispatcher.Dispatch(message)
	}
}
