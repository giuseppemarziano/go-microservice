package message

import (
	"context"
	"fmt"
	"github.com/palantir/stacktrace"
	"github.com/streadway/amqp"
	"go-microservice/infrastructure/message/messagebus"
	"log"
	"strings"
	"time"
)

type EventHandler interface {
	Handle(ctx context.Context, msg []byte) error
}

type Handler struct {
	bus      *Bus
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

const maxRetries = 3 // TODO add as env variable

func (h *Handler) DispatchToHandlers(msg messagebus.Message) {
	handler, exists := h.handlers[msg.RoutingKey]
	if !exists {
		log.Printf("No handler registered for message type %s", msg.RoutingKey)
		return
	}

	var err error
	for i := 0; i < maxRetries; i++ {
		err := handler.Handle(msg.Ctx, msg.Payload)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2) // TODO add env variable for wait time
	}

	if err != nil {
		log.Printf("Error handling message after %d attempts: %v", maxRetries, err)
		// TODO add logic to send the message to a dead-letter queue or other error handling mechanisms.
	}
}

func (h *Handler) StartListening(queueName string) error {
	messages, err := h.channel.Consume(
		queueName,
		"",
		false, // manual-ack (set to false)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return stacktrace.Propagate(err, "error consuming")
	}

	go func() {
		for d := range messages {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // Adjust the timeout as needed
			defer cancel()

			msg := messagebus.Message{
				RoutingKey: d.RoutingKey,
				Payload:    d.Body,
				Ctx:        ctx,
			}
			err := h.processMessage(msg)
			if err != nil {
				log.Printf("Error processing message: %v", err)
			} else {
				err := d.Ack(false)
				if err != nil {
					return
				}
			}
		}
	}()
	return nil
}

func (h *Handler) processMessage(msg messagebus.Message) error {
	handler, exists := h.handlers[msg.RoutingKey]
	if !exists {
		return fmt.Errorf("no handler registered for message type %s", msg.RoutingKey)
	}

	return handler.Handle(msg.Ctx, msg.Payload)
}

func (h *Handler) Close() error {
	var errs []string

	if h.channel != nil {
		err := h.channel.Close()
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if h.conn != nil {
		err := h.conn.Close()
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while closing: %s", strings.Join(errs, "; "))
	}
	return nil
}
