package f1tel

//
// const defaultSessionID = uint64(0)
//
// type _sessionManager struct {
// 	currSessionID uint64
// 	prevSessionID uint64
//
// 	eventStream MessagePublisher
// }
//
// func newSessionStateManager(eventStream MessagePublisher) *_sessionManager {
// 	return &_sessionManager{
// 		currSessionID: defaultSessionID,
// 		prevSessionID: defaultSessionID,
// 		eventStream:   eventStream,
// 	}
// }
//
// func (m *_sessionManager) StartSessionIfNotExist(sessionID uint64) {
// 	m.setCurrentSessionID(sessionID)
//
// 	// newSessionShouldBeCreated := m.prevSessionID == defaultSessionID && m.currSessionID != defaultSessionID
// 	//
// 	// if newSessionShouldBeCreated {
// 	// 	sessionStarted := telemetry.New(telemetry.SessionStartedMessageType, m.currSessionID, &telemetry.EmptyPayload)
// 	//
// 	// 	if err := m.eventStream.Publish(&sessionStarted, streams.SessionStartedSubject); err != nil {
// 	// 		log.Printf("failed to publish message: %s", err)
// 	// 	}
// 	// }
// }
//
// func (m *_sessionManager) setCurrentSessionID(value uint64) {
// 	m.prevSessionID = m.currSessionID
// 	m.currSessionID = value
// }
//
// func (m *_sessionManager) FinishSession() {
// 	// sessionFinished := telemetry.New(telemetry.SessionFinishedMessageType, m.currSessionID, &telemetry.EmptyPayload)
// 	//
// 	// if err := m.eventStream.Publish(&sessionFinished, streams.SessionFinishedSubject); err != nil {
// 	// 	log.Printf("failed to publish message: %s", err)
// 	// }
//
// 	m.setCurrentSessionID(defaultSessionID)
// }
