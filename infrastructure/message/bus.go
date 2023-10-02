package message

import (
	"go-microservice/infrastructure/message/messagebus"
	"sync"
)

type SubscriberFunc func(msg messagebus.Message)

type MessageBus struct {
	subscribers map[string][]SubscriberFunc
	mu          sync.RWMutex
}

func NewMessageBus() *MessageBus {
	return &MessageBus{
		subscribers: make(map[string][]SubscriberFunc),
	}
}

func (mb *MessageBus) Subscribe(routingKey string, subscriber SubscriberFunc) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	mb.subscribers[routingKey] = append(mb.subscribers[routingKey], subscriber)
}

func (mb *MessageBus) Publish(msg messagebus.Message) {
	mb.mu.RLock()
	defer mb.mu.RUnlock()

	if subscribers, ok := mb.subscribers[msg.RoutingKey]; ok {
		for _, subscriber := range subscribers {
			// Call subscriber in a goroutine for non-blocking processing
			go subscriber(msg)
		}
	}
}
