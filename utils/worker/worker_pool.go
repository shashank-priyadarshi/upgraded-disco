package worker

import "github.com/shashank-priyadarshi/upgraded-disco/utils/pubsub"

type Worker interface {
	SetPubSub(pubsub.PubSub)
	Close() error
	Wait()
}

type Pool struct {
	Count                 int
	Publisher, Subscriber string
}

func (w *Pool) SetPubSub(sub pubsub.PubSub) {
	//TODO implement me
	return
}

func (w *Pool) Close() error {
	//TODO implement me
	return nil
}

func (w *Pool) Wait() {
	//TODO implement me
	return
}
