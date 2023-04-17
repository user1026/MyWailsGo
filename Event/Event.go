package Event

import (
	"changeme/Global"
	"context"
)

type Event struct {
	ctx context.Context
}

func NewEventList() *Event {
	return &Event{}
}
func (e *Event) Startup(ctx context.Context) {
	e.ctx = ctx
	Global.Global_ConText = &ctx
}
func (e *Event) GetEvent(data ...interface{}) {

}
func openFile(data ...interface{}) {

}
