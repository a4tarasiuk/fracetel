package streams

import (
	"fracetel/core/messages"
)

const (
	SessionStreamName      = "sessions"
	SessionStreamSubject   = "sessions.*"
	SessionStartedSubject  = "sessions.started"
	SessionFinishedSubject = "sessions.finished"
)

var MessageTypeSubjectMap = map[messages.MessageType]string{}
