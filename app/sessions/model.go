package sessions

import "time"

type UserSession struct {
	ExternalID string `json:"external_id" bson:"external_id"` // TODO: [db] Create an index in a collection

	StartedAt time.Time `json:"started_at" bson:"started_at"`

	FinishedAt *time.Time `json:"finished_at" bson:"finished_at"`
}
