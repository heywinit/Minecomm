package models

type MinecraftPacket struct {
	PacketID int32 //VarInt, different for each type of packet
	Data []byte
}