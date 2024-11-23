package packets

type EventDataCode string

const (
	sessionStartedCode  EventDataCode = "SSTA"
	sessionFinishedCode               = "SEND"
)

type Event struct {
	Code1 uint8
	Code2 uint8
	Code3 uint8
	Code4 uint8
}

func (e Event) CodeToString() EventDataCode {
	return EventDataCode(string(e.Code1) + string(e.Code2) + string(e.Code3) + string(e.Code4))
}

func (e Event) IsSessionStarted() bool {
	return e.CodeToString() == sessionStartedCode
}

func (e Event) IsSessionFinished() bool {
	return e.CodeToString() == sessionFinishedCode
}
