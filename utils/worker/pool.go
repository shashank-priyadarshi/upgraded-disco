package worker

import (
	"github.com/shashank-priyadarshi/upgraded-disco/utils/pubsub"
)

type Worker interface {
	SetPubSub(pubsub.PubSub)
	Close() error
}

type Pool struct {
	count int
}

func NewPool(count int) (*Pool, error) {
	if count == 0 {
		count = 5
	}

	return &Pool{
		count: count,
	}, nil
}

func (p *Pool) SetPubSub(sub pubsub.PubSub) {
	//TODO implement me
	return
}

func (p *Pool) Close() error {
	//TODO implement me
	return nil
}
