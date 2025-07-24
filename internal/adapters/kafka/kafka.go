package kafka

import (
	"awesomeProject2/internal/adapters/dto/messagies"
	"awesomeProject2/internal/core/usecases/create"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type Consumer struct {
	reader  *kafka.Reader
	creator *create.UseCase
}

func NewConsumer(reader *kafka.Reader, creator *create.UseCase) *Consumer {
	return &Consumer{reader: reader, creator: creator}
}

func (c *Consumer) Consume() {
	ctx := context.Background()
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			time.Sleep(time.Second)
			continue
		}
		var msg messagies.Create
		err = json.Unmarshal(m.Value, &msg)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			time.Sleep(time.Second)
			continue
		}

		switch msg.Type {
		case "REGISTER":
			command := create.Command{ID: msg.UserID, Email: msg.Email, FirstName: msg.Username}
			err := c.creator.Handle(ctx, command)
			if err != nil {
				log.Printf("Error registering user: %v", err)
				continue
			}
		default:
			time.Sleep(time.Second)
			continue
		}
	}
}
