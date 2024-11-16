package packet

const F1TotalCars = 22

const (
	MotionID              ID = 0
	SessionID                = 1
	LapDataID                = 2
	EventID                  = 3
	ParticipantsID           = 4
	CarSetupsID              = 5
	CarTelemetryID           = 6
	CarStatusID              = 7
	FinalClassificationID    = 8
	LobbyInfoID              = 9
	CarDamageID              = 10
	SessionHistoryID         = 11
)

var IDName = map[ID]string{
	MotionID:              "Motion",
	SessionID:             "Session",
	LapDataID:             "LapData",
	EventID:               "Event",
	ParticipantsID:        "Participants",
	CarSetupsID:           "CarSetups",
	CarTelemetryID:        "CarTelemetry",
	CarStatusID:           "CarStatus",
	FinalClassificationID: "FinalClassification",
	LobbyInfoID:           "LobbyInfo",
	CarDamageID:           "CarDamage",
	SessionHistoryID:      "SessionHistory",
}

var IDDescription = map[ID]string{
	MotionID:              "Contains all motion data for player's car",
	SessionID:             "Data about the session - track, time left",
	LapDataID:             "Data about all the lap times of cars in the session",
	EventID:               "Various notable events that happen during session",
	ParticipantsID:        "List of participants in the session",
	CarSetupsID:           "Packet detailing car setups for cars in the race",
	CarTelemetryID:        "Telemetry data for all cars",
	CarStatusID:           "Status data for all cars",
	FinalClassificationID: "Final classification confirmation at the end of a race",
	LobbyInfoID:           "Information about players in a multiplayer lobby",
	CarDamageID:           "Damage status for all cars",
	SessionHistoryID:      "Lap and tyre data for session",
}
