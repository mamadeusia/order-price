package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"preh/config"
	"preh/core"

	"github.com/streadway/amqp"
)

var (
	InitMq      = initMq
	Publish     = publish
	RunConsumer = runConsumer
	DeclareQ    = declareQ
)

var (
	channel      *amqp.Channel
	publishQueue *amqp.Queue
	consumeQueue *amqp.Queue
	conn         *amqp.Connection
)

type PricePair struct {
	Price int64  `json:"price"`
	Pair  string `json:"pair"`
}

func initMq() (err error) {
	url := config.GetRabbitMQConnectionString()
	conn, err = amqp.Dial(url)
	if err != nil {
		return err
	}
	channel, err = conn.Channel()
	return err
}

func publish(exchange, queue string, body []byte) error {
	return channel.Publish(exchange, queue, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
}

func runConsumer(ctx context.Context) {
	messages := make(chan amqp.Delivery)
	go consumerWorker(ctx, messages)
	messageChannel := getMessageSubscribeChannel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("rabbit :: context done met!")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			messages <- message
		}
	}
}

func consumerWorker(ctx context.Context, messages <-chan amqp.Delivery) {
	for msg := range messages {
		if err := cunsumePrice(msg); err != nil {
			if err := msg.Reject(false); err != nil {
				fmt.Println("rabbit :: can't Reject message", err)
			}
		} else {
			if err := msg.Ack(false); err != nil {
				fmt.Println("rabbit :: can't Ack message", err)
			}
		}
	}
}
func declareQ() {
	channel.QueueDeclare("price", true, false, false, false, nil)
}
func cunsumePrice(delivery amqp.Delivery) error {
	var newPricePair PricePair
	err := json.Unmarshal(delivery.Body, &newPricePair)

	if err != nil {
		fmt.Println("rabbit :: can't unmarshal data")
	}

	core.SetCurrentPairPrice(newPricePair.Pair, newPricePair.Price)
	return nil
}

func getMessageSubscribeChannel() <-chan amqp.Delivery {
	messageChannel, err := channel.Consume(
		"price",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("dd")
		// log.Fatal("MQ issue" + err.Error() + " for queue: " + string(q))
	}

	err = channel.Qos(
		1,
		0,
		true,
	)
	if err != nil {
		fmt.Println("dd")

		// log.Error("No qos limit ", err)
	}

	return messageChannel
}
