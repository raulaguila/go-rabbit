package main

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/lucasjones/reggen"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/raulaguila/go-rabbit/internal/domain"
	"github.com/raulaguila/go-rabbit/pkg/alert"
	"github.com/raulaguila/go-rabbit/pkg/rabbitmq"
)

const (
	numberLocations int = 5
	numberAntennas  int = 4
)

var tags []*domain.TagReaded

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err.Error())
	}
}

func consumerQueueTags(rabbit *rabbitmq.Rabbitmq) {
	out := make(chan amqp.Delivery)
	go rabbit.Consume(".tags.", out)

	for msg := range out {
		tag := &domain.TagReaded{}
		if err := json.Unmarshal(msg.Body, tag); err != nil {
			log.Printf("Error unmarshal tag: %v\n", err.Error())
		} else {
			tag.ReaderAt = time.Now()
			tags = append(tags, tag)
		}
		msg.Ack(false)
	}
}

func checkSliceTags(rabbit *rabbitmq.Rabbitmq) {
	for {
		if len(tags) == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		tag := tags[0]
		tag.Location = fmt.Sprintf("ARMAZEM %02v", rand.Intn(numberLocations)+1)
		body, err := json.Marshal(tag)
		alert.Error(err)

		tags = tags[1:]
		alert.Error(rabbit.Publish("amq.topic", ".tags.detected", body))
	}
}

func readerSimulator(rabbit *rabbitmq.Rabbitmq, sleep time.Duration) {
	randString := func(n int) string {
		aux := ""
		for i := 0; i < 12; i++ {
			aux = fmt.Sprintf("[0-9a-fA-F]{2} %v", aux)
		}

		str, err := reggen.Generate(fmt.Sprintf("^%v$", aux), 10)
		alert.Error(err)
		return strings.TrimSpace(str)
	}

	randMac := func() string {
		buf := make([]byte, 6)
		_, err := crand.Read(buf)
		if err != nil {
			fmt.Println("error:", err)
			return ""
		}

		buf[0] |= 2
		return fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	}

	for {
		tag := domain.TagReaded{
			Mac:     randMac(),
			Tag:     randString(16),
			Antenna: uint(rand.Intn(numberAntennas) + 1),
		}

		body, _ := json.Marshal(tag)
		alert.Error(rabbit.Publish("amq.topic", ".tags.reader", body))
		time.Sleep(sleep)
	}
}

func main() {
	rabbit := &rabbitmq.Rabbitmq{}

	alert.Error(rabbit.OpenChannel())
	defer rabbit.CloseChannel()

	go consumerQueueTags(rabbit)
	go readerSimulator(rabbit, 200*time.Millisecond)
	checkSliceTags(rabbit)
}
