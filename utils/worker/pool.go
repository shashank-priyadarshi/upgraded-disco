package worker

import (
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"log"
	"sync"

	"github.com/shashank-priyadarshi/upgraded-disco/utils/pubsub"
)

type Worker interface {
	SetPubSub(pubsub.PubSub)
	Close() error
}

type Pool struct {
	broker                 pubsub.PubSub
	count                  int
	wg                     sync.WaitGroup
	mu                     sync.Locker
	subscribeTo, publishTo string
	subscriberIDs          map[string][]string
	log                    logger.Logger
}

func NewPool(count int, subscribeTo, publishTo string, log logger.Logger) (*Pool, error) {
	if count == 0 {
		count = 5
	}

	subscriberIDs := make(map[string][]string)
	subscriberIDs[subscribeTo] = []string{}
	subscriberIDs[publishTo] = []string{}

	return &Pool{
		count:         count,
		subscribeTo:   subscribeTo,
		publishTo:     publishTo,
		subscriberIDs: subscriberIDs,
		log:           log,
		mu:            &sync.Mutex{},
	}, nil
}

func (p *Pool) SetPubSub(sub pubsub.PubSub) {
	p.broker = sub

	for i := 0; i < p.count; i++ {
		p.wg.Add(1)
		go p.workerRoutine(i)
	}
}

func (p *Pool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	for topic, subscribers := range p.subscriberIDs {
		for _, subscriber := range subscribers {
			if err := p.broker.Unsubscribe(topic, subscriber); err != nil {
				p.log.Errorf(fmt.Sprintf("Error unsubcribing from topic %s for subscriber with ID %s: %v", topic, subscriber, err))
			}
		}
	}

	p.wg.Wait() // Wait for all worker routines to finish

	return nil
}

func (p *Pool) workerRoutine(workerID int) {
	p.log.Infof("Worker routine %d processing jobs received on channel %s", workerID, p.subscribeTo)

	subscriberID := p.broker.Subscribe(p.subscribeTo, func(message interface{}) {
		log.Printf("Worker-%d: Processing job - %v\n", workerID, message)

		result := fmt.Sprintf("Result of job %v processed by Worker-%d", message, workerID)
		p.broker.Publish(p.publishTo, result)
	})

	p.subscriberIDs[p.subscribeTo] = append(p.subscriberIDs[p.subscribeTo], subscriberID)

	p.mu.Lock()
	defer p.wg.Done()
	defer p.mu.Unlock()
	defer p.broker.Unsubscribe(p.subscribeTo, subscriberID)
}
