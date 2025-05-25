package telemetry

import "time"

type FinalClassification struct {
	SessionID       uint64
	FrameIdentifier uint32
	OccurredAt      time.Time

	FinishingPosition int

	StartingPosition int

	BestLapTimeMs float32
}
