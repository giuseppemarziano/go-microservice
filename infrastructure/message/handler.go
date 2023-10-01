package message

import (
	"context"
	"github.com/streadway/amqp"
	"log"
)

type EventHandler interface {
	Handle(ctx context.Context, msg interface{}) error
}

type Handler struct {
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
		conn:     conn,
		channel:  channel,
		handlers: make(map[string]EventHandler),
	}, nil
}

func (h *Handler) RegisterEventHandler(eventType string, handler EventHandler) {
	h.handlers[eventType] = handler
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
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			handler, exists := h.handlers[d.RoutingKey]
			if !exists {
				log.Printf("No handler registered for message type %s", d.RoutingKey)
				continue
			}

			err := handler.Handle(context.Background(), d.Body)
			if err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	<-forever
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
