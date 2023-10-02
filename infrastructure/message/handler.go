package message

import (
	"context"
	"github.com/palantir/stacktrace"
	"github.com/streadway/amqp"
	"go-microservice/infrastructure/message/messagebus"
	"log"
)

type EventHandler interface {
	Handle(ctx context.Context, msg []byte) error
}

type Handler struct {
	bus      *MessageBus
	conn     *amqp.Connection
	channel  *amqp.Channel
	handlers map[string]EventHandler
}

func NewHandler(connectionString string) (*Handler, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Handler{
		bus:      NewMessageBus(),
		conn:     conn,
		channel:  channel,
		handlers: make(map[string]EventHandler),
	}, nil
}

func (h *Handler) RegisterEventHandler(eventType string, handler EventHandler) {
	h.bus.Subscribe(eventType, h.DispatchToHandlers)
	h.handlers[eventType] = handler
}

func (h *Handler) DispatchToHandlers(msg messagebus.Message) {
	handler, exists := h.handlers[msg.RoutingKey]
	if !exists {
		log.Printf("No handler registered for message type %s", msg.RoutingKey)
		return
	}

	err := handler.Handle(context.Background(), msg.Payload)
	if err != nil {
		log.Printf("Error handling message: %v", err)
	}
}

func (h *Handler) StartListening(queueName string) error {
	messages, err := h.channel.Consume(
		queueName,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return stacktrace.Propagate(err, "error consuming ")
	}

	go func() {
		for d := range messages {
			msg := messagebus.Message{
				RoutingKey: d.RoutingKey,
				Payload:    d.Body,
			}
			h.bus.Publish(msg)
		}
	}()
	return nil
}

func (h *Handler) Close() error {
	if h.channel != nil {
		err := h.channel.Close()
		if err != nil {
			return err
		}
	}
	if h.conn != nil {
		return h.conn.Close()
	}
	return nil
}
