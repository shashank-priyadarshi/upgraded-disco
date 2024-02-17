package pubsub

type PubSub interface {
	Publish(string, string) error
	Subscribe(string, string, func(message string))
	Unsubscribe(string, string) error
}

type Broker struct{}

func NewBroker() (*Broker, error) {
	return &Broker{}, nil
}

func (b *Broker) Publish(topic string, message string) error {
	//TODO implement me
	return nil
}

func (b *Broker) Subscribe(topic string, subscriberID string, callback func(message string)) {
	//TODO implement me
	return
}

func (b *Broker) Unsubscribe(topic string, subscriberID string) error {
	//TODO implement me
	return nil
}
