package message

import (
	"go-microservice/infrastructure/message/messagebus"
	"sync"
)

type SubscriberFunc func(msg messagebus.Message)

type Bus struct {
	subscribers map[string][]SubscriberFunc
	mu          sync.RWMutex
	tasks       chan messagebus.Message
}

const numWorkers = 10 // TODO add env variable

func NewMessageBus() *Bus {
	mb := &Bus{
		subscribers: make(map[string][]SubscriberFunc),
		tasks:       make(chan messagebus.Message, 100), // Buffered channel; adjust size as needed
	}

	for i := 0; i < numWorkers; i++ {
		go mb.worker()
	}

	return mb
}

func (mb *Bus) worker() {
	for msg := range mb.tasks {
		if subscribers, ok := mb.subscribers[msg.RoutingKey]; ok {
			for _, subscriber := range subscribers {
				subscriber(msg)
			}
		}
	}
}

func (mb *Bus) Subscribe(routingKey string, subscriber SubscriberFunc) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	mb.subscribers[routingKey] = append(mb.subscribers[routingKey], subscriber)
}

func (mb *Bus) Publish(msg messagebus.Message) {
	mb.mu.RLock()
	defer mb.mu.RUnlock()

	mb.tasks <- msg
}
