package packets

type Header struct {
	PacketFormat uint16

	GameMajorVersion uint8
	GameMinorVersion uint8

	PacketVersion uint8
	PacketID      uint8

	SessionUID  uint64
	SessionTime float32

	FrameIdentifier uint32

	PlayerCarIdx          uint8
	SecondaryPlayerCarIdx uint8
}
