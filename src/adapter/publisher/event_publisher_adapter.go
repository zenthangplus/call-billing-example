package publisher

import (
	"gitlab.com/golibs-starter/golib/pubsub"
)

type EventPublisherAdapter struct {
}

func NewEventPublisherAdapter() *EventPublisherAdapter {
	return &EventPublisherAdapter{}
}

func (e EventPublisherAdapter) Publish(event pubsub.Event) {
	pubsub.Publish(event)
}
