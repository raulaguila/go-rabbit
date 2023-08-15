package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/raulaguila/go-rabbit/internal/tag/entity"
)

func getUri() string {
	if os.Getenv("RABBIT_USE") == "INT" {
		return fmt.Sprintf("amqp://%v:%v@%v:%v", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASS"), os.Getenv("RABBIT_INT_HOST"), os.Getenv("RABBIT_INT_PORT"))
	}
	return fmt.Sprintf("amqp://%v:%v@%v:%v", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASS"), os.Getenv("RABBIT_EXT_HOST"), os.Getenv("RABBIT_EXT_PORT"))
}

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial(getUri())
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume("tags", "go-consumer", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, tag entity.Tag) error {
	body, err := json.Marshal(tag)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	return ch.PublishWithContext(ctx, "amq.direct", "", false, false, amqp.Publishing{
		ContentType:     "application/json",
		ContentEncoding: "utf-8",
		Timestamp:       time.Now(),
		Body:            body,
	})
}
