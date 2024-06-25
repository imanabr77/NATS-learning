package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

const server = "nats://127.0.0.1:4222"

type Message struct {
	ID string
}

func main() {
	c, err := nats.Connect(server)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(c.ConnectedAddr())
	log.Print(c.DiscoveredServers())

	ec, err := nats.NewEncodedConn(c, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan struct{})

	if _, err := ec.Subscribe("hello-world", func(m *Message) {
		log.Print(m)
		close(ch)

	}); err != nil {
		log.Fatal(err)
	}

	if err := ec.Publish("hello-world", Message{ID: "1381"}); err != nil {
		log.Fatal(err)

	}

	<-ch
}
