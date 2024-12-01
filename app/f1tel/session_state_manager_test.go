package f1tel

import (
	"testing"

	"fracetel/core/messages"
	"fracetel/core/streams"
	"github.com/stretchr/testify/assert"
)

func Test_sessionManager_setCurrentSessionID(t *testing.T) {
	publisher := newInMemoryMessagePublisher()
	manager := newSessionStateManager(publisher)

	assert.Equal(t, defaultSessionID, manager.prevSessionID, "invalid initial value for prevSessionID")
	assert.Equal(t, defaultSessionID, manager.currSessionID, "invalid initial value for currSessionID")

	nextSessionID1 := uint64(15)

	manager.setCurrentSessionID(nextSessionID1)

	assert.Equal(t, defaultSessionID, manager.prevSessionID, "invalid initial value for prevSessionID")
	assert.Equal(t, nextSessionID1, manager.currSessionID, "currSessionID is not updated")

	nextSessionID2 := uint64(56)

	manager.setCurrentSessionID(nextSessionID2)

	assert.Equal(t, nextSessionID1, manager.prevSessionID, "currSessionID is not updated")
	assert.Equal(t, nextSessionID2, manager.currSessionID, "currSessionID is not updated")
}

func Test_sessionManager_WhenNewSessionCreated(t *testing.T) {
	publisher := newInMemoryMessagePublisher()
	manager := newSessionStateManager(publisher)

	assert.Equal(t, defaultSessionID, manager.prevSessionID, "invalid initial value for prevSessionID")

	nextSessionID := uint64(63142)

	manager.StartSessionIfNotExist(nextSessionID)

	publishedMessages, ok := publisher.messages[streams.SessionStartedSubject]
	assert.True(t, ok)

	expectedTotalPublishedMessages := 1
	assert.Equal(t, expectedTotalPublishedMessages, len(publishedMessages))

	msg := publishedMessages[0]

	assert.Equal(t, messages.SessionStartedMessageType, msg.Type, "Invalid message type")
	assert.Equal(t, nextSessionID, msg.Header.SessionID)
	assert.Equal(t, &messages.EmptyPayload, msg.Payload, "This message must be without payload")
	assert.False(t, msg.Header.OccurredAt.IsZero(), "DateTime must be set")
}

func Test_sessionManager_FinishSession(t *testing.T) {
	publisher := newInMemoryMessagePublisher()
	manager := newSessionStateManager(publisher)

	sessionID := uint64(45234)

	manager.setCurrentSessionID(sessionID)

	manager.FinishSession()

	publishedMessages, ok := publisher.messages[streams.SessionFinishedSubject]
	assert.True(t, ok, "There are no messages published with expected subject")

	expectedTotalPublishedMessages := 1
	assert.Equal(t, expectedTotalPublishedMessages, len(publishedMessages))

	msg := publishedMessages[0]

	assert.Equal(t, messages.SessionFinishedMessageType, msg.Type, "Invalid message type")
	assert.Equal(t, sessionID, msg.Header.SessionID)
	assert.Equal(t, &messages.EmptyPayload, msg.Payload, "This message must be without payload")
	assert.False(t, msg.Header.OccurredAt.IsZero(), "DateTime must be set")
}
