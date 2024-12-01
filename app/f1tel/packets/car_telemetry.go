package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/core/telemetry"
)

type carTelemetry struct {
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

	BrakesTemperature [4]uint16

	TyresSurfaceTemperature [4]uint8

	TyresInnerTemperature [4]uint8

	EngineTemperature uint16

	TyresPressure [4]float32

	SurfaceType [4]uint8
}

func (ct carTelemetry) ToTelemetryMessagePayload() telemetry.CarTelemetry {
	return telemetry.CarTelemetry{
		Speed:     int(ct.Speed),
		Throttle:  ct.Throttle,
		Steer:     ct.Steer,
		Brake:     ct.Brake,
		EngineRPM: int(ct.EngineRPM),
		DRS:       ct.DRS,
		TyreSurfaceTemperature: []int{
			int(ct.TyresSurfaceTemperature[0]),
			int(ct.TyresSurfaceTemperature[1]),
			int(ct.TyresSurfaceTemperature[2]),
			int(ct.TyresSurfaceTemperature[3]),
		},
		TyreInnerTemperature: []int{
			int(ct.TyresInnerTemperature[0]),
			int(ct.TyresInnerTemperature[1]),
			int(ct.TyresInnerTemperature[2]),
			int(ct.TyresInnerTemperature[3]),
		},
	}
}

type CarTelemetryPacketParser struct{}

func (p CarTelemetryPacketParser) ToTelemetryMessage(header *Header, rawPacket RawPacket) (
	*telemetry.Message,
	error,
) {
	telemetries := make([]carTelemetry, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &telemetries)

	if err != nil {
		log.Printf("Error during reading CarTelemetry: %s", err)
	}

	telemetryPayload := telemetries[header.PlayerCarIdx].ToTelemetryMessagePayload()

	msg := telemetry.NewMessage(
		telemetry.CarTelemetryMessageType,
		header.SessionUID,
		header.FrameIdentifier,
		&telemetryPayload,
	)

	log.Printf("Car Telemetry: %+v\n", msg.Payload)

	return &msg, nil
}
