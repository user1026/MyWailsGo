package Event

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Event struct {
	ctx context.Context
}

func NewEventList() *Event {
	return &Event{}
}
func (e *Event) Startup(ctx context.Context) {
	e.ctx = ctx
	runtime.EventsOn(e.ctx, "openFile", openFile)
}
func (e *Event) GetEvent(data ...interface{}) {
	runtime.EventsEmit(e.ctx, "openFile", "123123")
}
func openFile(data ...interface{}) {

}
