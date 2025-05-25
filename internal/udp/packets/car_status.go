package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/internal/messaging"
	"fracetel/pkg/telemetry"
)

type carStatus struct {
	TractionControl uint8
	AntiLockBrakes  uint8

	FuelMix uint8

	FrontBrakeBias uint8

	PutLimiterStatus uint8

	FuelInTank        float32
	FuelCapacity      float32
	FuelRemainingLaps float32

	MaxRPM  uint16
	IdleRPM uint16

	MaxGears uint8

	DRSAllowed            uint8
	DRSActivationDistance uint16

	ActualTyreCompound uint8
	VisualTyreCompound uint8

	TyresAgeLaps uint8

	VehicleFIAFlags int8

	ERSStoreEnergy float32
	ERSDeployMode  uint8

	ERSHarvestedThisLapMGUK float32
	ERSHarvestedThisLapMGUH float32
	ERSDeployedThisLap      float32

	NetworkPaused uint8
}

func (cs carStatus) ToTelemetryMessagePayload() telemetry.CarStatus {
	return telemetry.CarStatus{
		TractionControl:         int(cs.TractionControl),
		AntiLockBrakes:          int(cs.AntiLockBrakes),
		FuelMix:                 int(cs.FuelMix),
		FrontBrakeBias:          int(cs.FrontBrakeBias),
		PutLimiterStatus:        int(cs.PutLimiterStatus),
		FuelInTank:              cs.FuelInTank,
		FuelCapacity:            cs.FuelCapacity,
		FuelRemainingLaps:       cs.FuelRemainingLaps,
		MaxRPM:                  int(cs.MaxRPM),
		IdleRPM:                 int(cs.IdleRPM),
		MaxGears:                int(cs.MaxGears),
		DRSAllowed:              int(cs.DRSAllowed),
		DRSActivationDistance:   int(cs.DRSActivationDistance),
		ActualTyreCompound:      int(cs.ActualTyreCompound),
		VisualTyreCompound:      int(cs.VisualTyreCompound),
		TyresAgeLaps:            int(cs.TyresAgeLaps),
		VehicleFIAFlags:         int(cs.VehicleFIAFlags),
		ERSStoreEnergy:          cs.ERSStoreEnergy,
		ERSDeployMode:           int(cs.ERSDeployMode),
		ERSHarvestedThisLapMGUK: cs.ERSHarvestedThisLapMGUK,
		ERSHarvestedThisLapMGUH: cs.ERSHarvestedThisLapMGUH,
		ERSDeployedThisLap:      cs.ERSDeployedThisLap,
	}
}

type CarStatusParser struct{}

func (p CarStatusParser) ToTelemetryMessage(header *Header, rawPacket RawPacket) (
	*messaging.Message,
	error,
) {
	carStatusPackets := make([]carStatus, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &carStatusPackets)

	if err != nil {
		log.Printf("Error during reading CarStatus: %s", err)
	}

	carStatusPayload := carStatusPackets[header.PlayerCarIdx].ToTelemetryMessagePayload()

	msg := messaging.NewMessage(
		messaging.CarStatusMessageType,
		header.SessionUID,
		header.FrameIdentifier,
		&carStatusPayload,
	)

	log.Printf("Car Status: %+v\n", msg.Payload)

	return &msg, nil
}
