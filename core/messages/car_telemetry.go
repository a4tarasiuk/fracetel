package messages

type CarTelemetry struct {
	Speed uint16 `json:"speed"`

	Throttle float32 `json:"throttle"`
	Steer    float32 `json:"steer"`
	Brake    float32 `json:"brake"`

	EngineRPM uint16 `json:"engine_rpm"`

	DRS uint8 `json:"drs"`

	// INFO: Consider creating an object that represents 4 wheels with any values
	TyreSurfaceTemperature []int `json:"tyre_surface_temperature"`

	TyreInnerTemperature []int `json:"tyre_inner_temperature"`
}
