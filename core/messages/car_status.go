package messages

type CarStatus struct {
	TractionControl int
	AntiLockBrakes  int

	FuelMix int

	FrontBrakeBias int

	PutLimiterStatus int

	FuelInTank        float32
	FuelCapacity      float32
	FuelRemainingLaps float32

	MaxRPM  int
	IdleRPM int

	MaxGears int

	DRSAllowed            int
	DRSActivationDistance int

	ActualTyreCompound int
	VisualTyreCompound int

	TyresAgeLaps int

	VehicleFIAFlags int

	ERSStoreEnergy float32
	ERSDeployMode  int

	ERSHarvestedThisLapMGUK float32
	ERSHarvestedThisLapMGUH float32
	ERSDeployedThisLap      float32
}
