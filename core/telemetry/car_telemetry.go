package telemetry

type CarTelemetry struct {
	Speed int `json:"speed"`

	Throttle float32 `json:"throttle"`
	Steer    float32 `json:"steer"`
	Brake    float32 `json:"brake"`

	EngineRPM int `json:"engine_rpm"`

	DRS byte `json:"drs"`

	// INFO: Consider creating an object that represents 4 wheels with any values
	TyreSurfaceTemperature []int `json:"tyre_surface_temperature"`

	TyreInnerTemperature []int `json:"tyre_inner_temperature"`
}
