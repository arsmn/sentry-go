package sentry

import (
	"context"
	"time"
)

func Init(options ClientOptions) error {
	hub := CurrentHub()
	client, err := NewClient(options)
	if err != nil {
		return err
	}
	hub.BindClient(client)
	return nil
}

func AddBreadcrumb(breadcrumb *Breadcrumb) {
	hub := CurrentHub()
	hub.AddBreadcrumb(breadcrumb, nil)
}

// CaptureMessage captures an arbitrary message.
func CaptureMessage(message string) *EventID {
	hub := CurrentHub()
	return hub.CaptureMessage(message, nil)
}

func CaptureException(exception error) *EventID {
	hub := CurrentHub()
	return hub.CaptureException(exception, &EventHint{OriginalException: exception})
}

// CaptureEvent captures an event on the currently active client if any.
//
// The event must already be assembled. Typically code would instead use
// the utility methods like `CaptureException`. The return value is the
// event ID. In case Sentry is disabled or event was dropped, the return value will be nil.
func CaptureEvent(event *Event) *EventID {
	hub := CurrentHub()
	return hub.CaptureEvent(event, nil)
}

func Recover() {
	if err := recover(); err != nil {
		hub := CurrentHub()
		hub.Recover(err, &EventHint{RecoveredException: err})
	}
}

func RecoverWithContext(ctx context.Context) {
	if err := recover(); err != nil {
		var hub *Hub

		if HasHubOnContext(ctx) {
			hub = GetHubFromContext(ctx)
		} else {
			hub = CurrentHub()
		}

		hub.RecoverWithContext(ctx, err, &EventHint{RecoveredException: err})
	}
}

// TODO: Or maybe just `Recover(true)`? It may be too generic though
// func RecoverAndPanic() {
// 	if err := recover(); err != nil {
// 		Recover()
// 		panic(err)
// 	}
// }

func WithScope(f func(scope *Scope)) {
	hub := CurrentHub()
	hub.WithScope(f)
}

func ConfigureScope(f func(scope *Scope)) {
	hub := CurrentHub()
	hub.ConfigureScope(f)
}

func PushScope() {
	hub := CurrentHub()
	hub.PushScope()
}
func PopScope() {
	hub := CurrentHub()
	hub.PopScope()
}

func Flush(timeout time.Duration) bool {
	hub := CurrentHub()
	return hub.Flush(timeout)
}

func LastEventID() EventID {
	hub := CurrentHub()
	return hub.LastEventID()
}
