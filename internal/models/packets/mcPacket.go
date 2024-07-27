package packets

type MinecraftPacket struct {
	Length   int32  `mc:"length" json:"length"`     //VarInt, length of the packet
	PacketID int32  `mc:"packetID" json:"packetID"` //VarInt, different for each type of packet
	Data     []byte `mc:"data" json:"data"`
}
