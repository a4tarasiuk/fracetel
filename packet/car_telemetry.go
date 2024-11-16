package packet

type CarTelemetry struct {
	Speed uint16

	Throttle float32
	Steer    float32
	Brake    float32

	Clutch uint8
	Gear   int8

	EngineRPM uint16

	DRS uint8

	RevLightsPercent  uint8
	RevLightsBitValue uint16

	BrakesTemperatureRL uint16
	BrakesTemperatureRR uint16
	BrakesTemperatureFL uint16
	BrakesTemperatureFR uint16

	TyresSurfaceTemperatureRL uint8
	TyresSurfaceTemperatureRR uint8
	TyresSurfaceTemperatureFL uint8
	TyresSurfaceTemperatureFR uint8

	TyresInnerTemperatureRL uint8
	TyresInnerTemperatureRR uint8
	TyresInnerTemperatureFL uint8
	TyresInnerTemperatureFR uint8

	EngineTemperature uint16

	TyresPressureRL float32
	TyresPressureRR float32
	TyresPressureFL float32
	TyresPressureFR float32

	SurfaceType1 uint8
	SurfaceType2 uint8
	SurfaceType3 uint8
	SurfaceType4 uint8
}
