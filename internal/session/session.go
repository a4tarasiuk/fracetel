package session

import "time"

type Session struct {
	ID string

	FastestLapTimeMs int `db:"fastest_lap_time_ms"`

	FastestLapSector1Ms int `db:"fastest_lap_sector_1_ms"`
	FastestLapSector2Ms int `db:"fastest_lap_sector_2_ms"`
	FastestLapSector3Ms int `db:"fastest_lap_sector_3_ms"`

	FastestLapNumber uint8 `db:"fastest_lap_number"`

	TotalLaps uint8 `db:"total_laps"`
	TrackID   uint8 `db:"track_id"`

	Weather int `db:"weather"`

	Duration int `db:"duration"`

	StartingPosition  uint8 `db:"starting_position"`
	FinishingPosition uint8 `db:"finishing_position"`

	CreatedAt time.Time `db:"created_at"`
}
