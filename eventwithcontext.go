package dbr

import "context"

// EventReceiverWithContext gets events from dbr methods for logging purposes.
type EventReceiverWithContext interface {
	EventWithContext(ctx context.Context, eventName string)
	EventKvWithContext(ctx context.Context, eventName string, kvs map[string]string)
	EventErrWithContext(ctx context.Context, eventName string, err error) error
	EventErrKvWithContext(ctx context.Context, eventName string, err error, kvs map[string]string) error
	TimingWithContext(ctx context.Context, eventName string, nanoseconds int64)
	TimingKvWithContext(ctx context.Context, eventName string, nanoseconds int64, kvs map[string]string)
}

var nullReceiverWithContext = &NullEventReceiverWithContext{}

// NullEventReceiverWithContext is a sentinel EventReceiver.
// Use it if the caller doesn't supply one.
type NullEventReceiverWithContext struct{}

// EventWithContext receives a simple notification when various events occur.
func (n *NullEventReceiverWithContext) EventWithContext(ctx context.Context, eventName string) {}

// EventKvWithContext receives a notification when various events occur along with
// optional key/value data.
func (n *NullEventReceiverWithContext) EventKvWithContext(ctx context.Context, eventName string, kvs map[string]string) {
}

// EventErrWithContext receives a notification of an error if one occurs.
func (n *NullEventReceiverWithContext) EventErrWithContext(ctx context.Context, eventName string, err error) error {
	return err
}

// EventErrKvWithContext receives a notification of an error if one occurs along with
// optional key/value data.
func (n *NullEventReceiverWithContext) EventErrKvWithContext(ctx context.Context, eventName string, err error, kvs map[string]string) error {
	return err
}

// TimingWithContext receives the time an event took to happen.
func (n *NullEventReceiverWithContext) TimingWithContext(ctx context.Context, eventName string, nanoseconds int64) {
}

// TimingKvWithContext receives the time an event took to happen along with optional key/value data.
func (n *NullEventReceiverWithContext) TimingKvWithContext(ctx context.Context, eventName string, nanoseconds int64, kvs map[string]string) {
}
