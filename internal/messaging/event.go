package messaging

type (
	EventPayload interface{}

	Event interface {
		GetEventName() string

		GetEventPayload() EventPayload
	}
)
