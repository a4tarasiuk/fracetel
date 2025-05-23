package messaging

import "context"

type EventStream interface {
	Publish(ctx context.Context, topicName string, value Event) error
}
