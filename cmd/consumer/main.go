package main

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/raulaguila/go-rabbit/internal/tag/entity"
	"github.com/raulaguila/go-rabbit/pkg/rabbitmq"
)

const workers int = 2

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err.Error())
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func consume(number int) {
	ch, err := rabbitmq.OpenChannel()
	panicErr(err)
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	for msg := range out {
		msg.Ack(false)
		tag := entity.Tag{}
		err := json.Unmarshal(msg.Body, &tag)
		panicErr(err)

		// fmt.Printf("Consumer number %v: %v\n", number, tag)
	}
}

func main() {
	for i := 1; i < workers; i++ {
		go consume(i)
	}

	consume(workers)
}
