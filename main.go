package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
    // Connect to NATS server with credentials
    nc, err := nats.Connect("nats://sys:password@localhost:4222")
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

    // Create JetStream context
    js, err := nc.JetStream()
    if err != nil {
        log.Fatal(err)
    }

    // Create a stream
    streamConfig := &nats.StreamConfig{
        Name:     "mystream",
        Subjects: []string{"mystream.*"},
        Storage:  nats.FileStorage,
    }
    _, err = js.AddStream(streamConfig)
    if err != nil {
        log.Fatal(err)
    }

    // Publish a message to the stream
    _, err = js.Publish("mystream.test", []byte("Hello, NATS!"))
    if err != nil {
        log.Fatal(err)
    }

    // Create a consumer
    consumerConfig := &nats.ConsumerConfig{
        Durable:   "mystream-consumer",
        AckPolicy: nats.AckExplicitPolicy,
    }
    _, err = js.AddConsumer("mystream", consumerConfig)
    if err != nil {
        log.Fatal(err)
    }

    // Subscribe to the stream
    sub, err := js.PullSubscribe("mystream.test", "mystream-consumer")
    if err != nil {
        log.Fatal(err)
    }

    // Read messages from the stream
    msgs, err := sub.Fetch(1, nats.MaxWait(10*time.Second))
    if err != nil {
        log.Fatal(err)
    }

    for _, msg := range msgs {
        fmt.Printf("Received message: %s\n", string(msg.Data))
        msg.Ack()
    }
}