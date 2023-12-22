package events

import (
	"log"

	"github.com/nats-io/nats.go"
)

func ExampleJetStreamPublisher() *nats.Conn {
	nc, err := nats.Connect("nats1")
	if err != nil {
		log.Fatal(err)
	}

	// Use the JetStream context to produce and consumer messages
	// that have been persisted.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "FOO",
		Subjects: []string{"foo"},
	})

	return nc

}
