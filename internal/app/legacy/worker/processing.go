package worker

import (
	"fmt"
	"log"

	sessions2 "fracetel/internal/app/legacy/app/sessions"
	"github.com/google/uuid"
)

func processSessionMessage(
	sessionChan <-chan Session,
	userSessionService sessions2.UserSessionService,
) {
	sessionStateManager := sessions2.NewSessionStateManager()

	for session := range sessionChan {
		if !sessionStateManager.ShouldStartNewSession(session.SessionID) {
			continue
		}

		userSession := sessions2.UserSession{
			ID:         uuid.New().String(),
			ExternalID: session.SessionID,
			StartedAt:  session.OccurredAt,
			FinishedAt: nil,
			Type:       session.Type,
			TrackID:    session.TrackID,
			TotalLaps:  session.TotalLaps,
		}

		if err := userSessionService.StartSession(userSession); err != nil {
			log.Printf("Failed to start user session: %s", err)
		}
	}

	fmt.Println("Closed session processor")
}

func processFinalClassificationMessage(
	finalClassificationChan <-chan FinalClassification,
	userSessionService sessions2.UserSessionService,
) {
	for finalClassification := range finalClassificationChan {

		if err := userSessionService.FinishSession(finalClassification.SessionID); err != nil {
			log.Printf("Failed to finish user session: %s", err)
		}
	}
}
