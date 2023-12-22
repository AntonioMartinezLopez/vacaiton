package events

import (
	"backend/pkg/logger"

	"github.com/nats-io/nats.go"
)

func ExampleJetStreamGroupConsumer(stream string, subject string, group string, handler func(*nats.Msg)) {
	nc, err := nats.Connect("nats1")
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Use the JetStream context to produce and consumer messages
	// that have been persisted.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		logger.Fatal(err.Error())
	}

	js.AddStream(&nats.StreamConfig{
		Name:     stream,
		Subjects: []string{subject},
	})

	// ordered push consumer
	js.QueueSubscribe(subject, group, handler)
	//js.Subscribe("foo", handler, nats.OrderedConsumer())
}
