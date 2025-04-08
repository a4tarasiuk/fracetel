package sessions

import "time"

type UserSession struct {
	ID         string `bson:"id"`
	ExternalID string `bson:"external_id"`

	StartedAt  time.Time  `bson:"started_at"`
	FinishedAt *time.Time `bson:"finished_at"`

	Type int `bson:"type"`

	TrackID int `bson:"track_id"`

	TotalLaps int `bson:"total_laps"`
}
