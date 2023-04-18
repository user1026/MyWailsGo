package Event

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Event struct {
	ctx context.Context
}

func (event *Event) SetContext(ctx context.Context) {
	event.ctx = ctx
}
func (event *Event) JsToGo(EventName string, callBack func(data ...interface{})) {
	runtime.EventsOn(event.ctx, EventName, callBack)
}
func (event *Event) GotoJs(EventName string, data ...interface{}) {

	runtime.EventsEmit(event.ctx, EventName, data)
}

var EventList = new(Event)
