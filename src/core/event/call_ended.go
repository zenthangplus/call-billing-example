package event

import (
	"context"
	"github.com/zenthangplus/call-billing-example/src/core/entity"
	"gitlab.com/golibs-starter/golib/event"
	webEvent "gitlab.com/golibs-starter/golib/web/event"
)

func NewCallEnded(ctx context.Context, call *entity.Call) *CallEnded {
	return &CallEnded{
		AbstractEvent: webEvent.NewAbstractEvent(ctx, "CallEnded", event.WithPayload(call)),
	}
}

type CallEnded struct {
	*webEvent.AbstractEvent
}
