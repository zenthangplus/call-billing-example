package listener

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	event2 "github.com/zenthangplus/call-billing-example/src/core/event"
	"github.com/zenthangplus/call-billing-example/src/core/service"
	"gitlab.com/golibs-starter/golib/pubsub"
	"gitlab.com/golibs-starter/golib/web/log"
)

type BillingAggregation struct {
	service service.BillingService
}

func NewBillingAggregation(service service.BillingService) pubsub.Subscriber {
	return &BillingAggregation{service: service}
}

func (b BillingAggregation) Supports(event pubsub.Event) bool {
	_, ok := event.(*event2.CallEnded)
	return ok
}

func (b BillingAggregation) Handle(event pubsub.Event) {
	payload := event.Payload().(*entity.Call)
	if err := b.service.Aggregate(context.Background(), payload); err != nil {
		log.Errorf("Cannot aggregate billing for call [%d], error [%s]", payload.Id, err)
		return
	}
	log.Infof("Aggregate billing success for call [%d]", payload.Id)
}
