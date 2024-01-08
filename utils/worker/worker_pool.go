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
	panic("implement me")
}

func (w *Pool) Close() error {
	//TODO implement me
	panic("implement me")
}

func (w *Pool) Wait() {
	//TODO implement me
	panic("implement me")
}
