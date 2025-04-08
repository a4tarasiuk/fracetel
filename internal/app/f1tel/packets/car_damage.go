package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	"fracetel/internal/core/telemetry"
)

type carDamage struct {
	TyresWear [4]float32

	Tyres [4]uint8

	Brakes [4]uint8

	FrontLeftWing  uint8
	FrontRightWing uint8

	RearWing uint8

	Floor uint8

	Diffuser uint8

	SidePod uint8

	DRSFault uint8
	ERSFault uint8

	GearBox uint8

	Engine         uint8
	EngineMGUHWear uint8
	EngineESWear   uint8
	EngineCEWear   uint8
	EngineICEWear  uint8
	EngineMGUKWear uint8
	EngineTCWear   uint8
	EngineBlown    uint8
	EngineSeized   uint8
}

func (cd carDamage) ToTelemetryMessagePayload() telemetry.CarDamage {
	return telemetry.CarDamage{
		Tyres:     []int{int(cd.Tyres[0]), int(cd.Tyres[1]), int(cd.Tyres[2]), int(cd.Tyres[3])},
		TyresWear: []int{int(cd.TyresWear[0]), int(cd.TyresWear[1]), int(cd.TyresWear[2]), int(cd.TyresWear[3])},
		Brakes:    []int{int(cd.Brakes[0]), int(cd.Brakes[1]), int(cd.Brakes[2]), int(cd.Brakes[3])},
	}
}

type carDamageParser struct{}

func (p carDamageParser) ToTelemetryMessage(header *Header, rawPacket RawPacket) (
	*telemetry.Message,
	error,
) {
	carDamagePackets := make([]carDamage, F1TotalCars)

	buffer := bytes.NewBuffer(rawPacket[HeaderTotalBytes:])

	err := binary.Read(buffer, PacketByteOrder, &carDamagePackets)

	if err != nil {
		log.Printf("Error during reading CarDamage: %s", err)
	}

	carDamagePayload := carDamagePackets[header.PlayerCarIdx].ToTelemetryMessagePayload()

	msg := telemetry.NewMessage(
		telemetry.CarDamageMessageType,
		header.SessionUID,
		header.FrameIdentifier,
		&carDamagePayload,
	)

	log.Printf("Car Damage: %+v\n", msg.Payload)

	return &msg, nil
}
