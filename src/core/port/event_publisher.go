package port

import "gitlab.com/golibs-starter/golib/pubsub"

type EventPublisher interface {

	// Publish an event
	Publish(e pubsub.Event)
}
