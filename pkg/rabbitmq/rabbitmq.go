package rabbitmq

import (
	"context"
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	ch *amqp.Channel
}

func (r *Rabbitmq) getUri() string {
	if os.Getenv("RABBIT_USE") == "INT" {
		return fmt.Sprintf("amqp://%v:%v@%v:%v", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASS"), os.Getenv("RABBIT_INT_HOST"), os.Getenv("RABBIT_INT_PORT"))
	}
	return fmt.Sprintf("amqp://%v:%v@%v:%v", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASS"), os.Getenv("RABBIT_EXT_HOST"), os.Getenv("RABBIT_EXT_PORT"))
}

func (r *Rabbitmq) CloseChannel() error {
	return r.ch.Close()
}

func (r *Rabbitmq) OpenChannel() error {
	conn, err := amqp.Dial(r.getUri())
	if err != nil {
		return err
	}

	r.ch, err = conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (r *Rabbitmq) Consume(queue string, out chan amqp.Delivery) error {
	msgs, err := r.ch.Consume(queue, "go-consumer", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func (r *Rabbitmq) Publish(exchange string, topic string, body []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	return r.ch.PublishWithContext(ctx, exchange, topic, false, false, amqp.Publishing{
		ContentType:     "application/json",
		ContentEncoding: "utf-8",
		Timestamp:       time.Now(),
		Body:            body,
	})
}
