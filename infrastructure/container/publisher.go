package container

import "github.com/streadway/amqp"

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewPublisher(connectionString string) (*Publisher, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Publisher{conn: conn, channel: channel}, nil
}

func (p *Publisher) Publish(exchange, routingKey string, message []byte) error {
	err := p.channel.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	return err
}

func (p *Publisher) Close() error {
	if err := p.channel.Close(); err != nil {
		return err
	}
	if err := p.conn.Close(); err != nil {
		return err
	}
	return nil
}
