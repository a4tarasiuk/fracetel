package sessions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sessionManager_ShouldStartNewSession(t *testing.T) {
	manager := NewSessionStateManager()

	assert.Equal(t, defaultSessionID, manager.prevSessionID, "invalid initial value for prevSessionID")

	nextSessionID := "63142"

	assert.True(t, manager.ShouldStartNewSession(nextSessionID))

	anotherNewSessionID := "099101"

	assert.True(t, manager.ShouldStartNewSession(anotherNewSessionID), "new session must be detected")

	assert.False(
		t,
		manager.ShouldStartNewSession(anotherNewSessionID),
		"new session must not created, the same session ID is used",
	)
}
