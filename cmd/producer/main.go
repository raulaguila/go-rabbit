package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/raulaguila/go-rabbit/internal/tag/entity"
	"github.com/raulaguila/go-rabbit/pkg/rabbitmq"
)

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

func GenerateTag() entity.Tag {
	return entity.Tag{
		Tag:     uuid.New().String(),
		Reader:  uuid.New().String(),
		Antenna: rand.Uint64(),
	}
}

func publish(number int) {
	ch, err := rabbitmq.OpenChannel()
	panicErr(err)
	defer ch.Close()

	for {
		tag := GenerateTag()
		fmt.Printf("Producer number %v: %v\n", number, tag)
		panicErr(rabbitmq.Publish(ch, tag))
	}
}

func main() {
	for i := 1; i < 5; i++ {
		go publish(i)
		time.Sleep(time.Second)
	}

	publish(5)
}
