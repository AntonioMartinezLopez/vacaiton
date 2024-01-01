package events

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type NatsConnectorInterface interface {
	NatsJetStream() (nats.JetStreamContext, error)
	CreateStream(name string, subjects []string) (*nats.StreamInfo, error)
	PublishStream(subject string, message []byte) error
	PublishStreamAsync(subject string, message []byte) error
	PublishAsyncFinished() <-chan struct{}
	QueueSubscribe(stream string, subject string, group string, handler func(*nats.Msg)) (nats.Subscription, error)
}

type NatsConnector struct {
	Connection *nats.Conn
}

func NewNatsConnector(url string) (*NatsConnector, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("nats connect: %w", err)
	}

	return &NatsConnector{Connection: nc}, nil
}

func (nc *NatsConnector) NatsJetStream() (nats.JetStreamContext, error) {
	jsCtx, err := nc.Connection.JetStream()
	if err != nil {
		return nil, fmt.Errorf("jetstream: %w", err)
	}
	return jsCtx, nil
}

func (nc *NatsConnector) CreateStream(name string, subjects []string) (*nats.StreamInfo, error) {
	jsCtx, err := nc.NatsJetStream()

	if err != nil {
		return nil, err
	}

	stream, err := jsCtx.AddStream(&nats.StreamConfig{
		Name:      name,
		Subjects:  subjects,
		Retention: nats.InterestPolicy, // remove acked messages
		Discard:   nats.DiscardOld,     // when the stream is full, discard old messages
		MaxAge:    7 * 24 * time.Hour,  // max age of stored messages is 7 days
	})

	if err != nil {
		return nil, fmt.Errorf("add stream: %w", err)
	}

	return stream, nil

}

func (nc *NatsConnector) PublishStream(subject string, message []byte) error {

	jsCtx, err := nc.NatsJetStream()

	if err != nil {
		return err
	}

	_, publishErr := jsCtx.Publish(subject, message)

	if publishErr != nil {
		return fmt.Errorf("publish: %w", err)
	}

	return nil
}

func (nc *NatsConnector) PublishStreamAsync(subject string, message []byte) error {

	jsCtx, err := nc.NatsJetStream()

	if err != nil {
		return err
	}

	_, publishErr := jsCtx.PublishAsync(subject, message)

	if publishErr != nil {
		return fmt.Errorf("publish: %w", err)
	}

	return nil
}

func (nc *NatsConnector) PublishAsyncFinished() <-chan struct{} {
	jsCtx, err := nc.NatsJetStream()

	if err != nil {
		return nil
	}

	return jsCtx.PublishAsyncComplete()
}

func (nc *NatsConnector) QueueSubscribe(stream string, subject string, group string, handler func(*nats.Msg)) (*nats.Subscription, error) {

	jsCtx, err := nc.NatsJetStream()

	if err != nil {
		return nil, err
	}

	return jsCtx.QueueSubscribe(subject, group, handler)
}
