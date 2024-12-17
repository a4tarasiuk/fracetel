package sessions

const defaultSessionID = "0"

type sessionManager struct {
	currSessionID string
	prevSessionID string
}

func NewSessionStateManager() *sessionManager {
	return &sessionManager{
		currSessionID: defaultSessionID,
		prevSessionID: defaultSessionID,
	}
}

func (m *sessionManager) ShouldStartNewSession(sessionID string) bool {
	m.prevSessionID = m.currSessionID
	m.currSessionID = sessionID

	defaultSessionFinished := m.prevSessionID == defaultSessionID
	defaultSessionStarted := m.currSessionID != defaultSessionID
	newUserSessionStarted := m.currSessionID != m.prevSessionID

	shouldStartNewSession := (defaultSessionFinished && defaultSessionStarted) || newUserSessionStarted

	return shouldStartNewSession
}
