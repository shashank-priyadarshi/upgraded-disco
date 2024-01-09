package pubsub

type PubSub interface {
	Publish(topic string, message string) error
	Subscribe(topic string, subscriberID string, callback func(message string))
	Unsubscribe(topic string, subscriberID string) error
}

type Broker struct{}

func NewBroker() (broker Broker, err error) { return }

func (b Broker) Publish(topic string, message string) error {
	//TODO implement me
	return nil
}

func (b Broker) Subscribe(topic string, subscriberID string, callback func(message string)) {
	//TODO implement me
	return
}

func (b Broker) Unsubscribe(topic string, subscriberID string) error {
	//TODO implement me
	return nil
}
