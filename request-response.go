package main

import (
	"encoding/json"
	"log"
	"time"

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

	if _, err := c.Subscribe("hello-world", func(msg *nats.Msg) {
		if err := c.Publish(msg.Reply, []byte("Hello-world")); err != nil {
			log.Fatal(err)
		}
	}); err != nil {
		log.Fatal(err)
	}

	d, err := json.Marshal(Message{ID: "1381"})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Request("hello-world", d, 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(resp.Data))
	log.Print("Request Success")
}
