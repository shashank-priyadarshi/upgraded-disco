package pubsub

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type PubSub interface {
	Publish(string, string) error
	Subscribe(string, func(message interface{})) string
	Unsubscribe(string, string) error
}

type Broker struct {
	mu            sync.Locker
	topics        map[string][]chan interface{}
	subscribers   map[string]chan interface{}
	subscriberMux sync.Mutex
}

func NewBroker() (*Broker, error) {
	return &Broker{
		mu:          &sync.Mutex{},
		topics:      make(map[string][]chan interface{}),
		subscribers: make(map[string]chan interface{}),
	}, nil
}

func (b *Broker) Publish(topic string, message string) error {
	b.exists(topic)

	for _, c := range b.topics[topic] {
		c <- message
	}

	return nil
}

func (b *Broker) Subscribe(topic string, callback func(message interface{})) string {
	b.mu.Lock()
	defer b.mu.Unlock()

	subscriberChannel := make(chan interface{})
	subscriberID := uuid.New().String()

	b.exists(topic)
	b.topics[topic] = append(b.topics[topic], subscriberChannel)
	b.subscribers[subscriberID] = subscriberChannel

	go func() {
		for {
			message := <-subscriberChannel
			callback(message)
		}
	}()

	return subscriberID
}

func (b *Broker) Unsubscribe(topic string, subscriberID string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	subscriberChannel, ok := b.subscribers[subscriberID]
	if !ok {
		return errors.New("subscriber isn't found")
	}

	var newChannels []chan interface{}
	for _, c := range b.topics[topic] {
		if c != subscriberChannel {
			newChannels = append(newChannels, c)
		}
	}
	b.topics[topic] = newChannels

	delete(b.subscribers, subscriberID)

	close(subscriberChannel)

	return nil
}

func (b *Broker) exists(topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.topics[topic]; !ok {
		b.topics[topic] = make([]chan interface{}, 0)
	}

	return
}
