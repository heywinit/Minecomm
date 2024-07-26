package models

//HandShakePacket, serverbound
type HandShakePacket struct {
	MinecraftPacket

	ProtocolVersion int32 //VarInt
	ServerAddress   string //max: 255
	ServerPort      uint16
	NextState       int8 //VarInt enum, 1 for status, 2 for login, 3 for transfer 
}