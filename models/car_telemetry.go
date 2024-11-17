package models

type CarTelemetry struct {
	Speed uint16

	Throttle float32
	Steer    float32
	Brake    float32

	EngineRPM uint16

	DRS uint8

	// INFO: Consider creating an object that represents 4 wheels with any values
	TyreSurfaceTemperature []uint8

	TyreInnerTemperature []uint8
}
