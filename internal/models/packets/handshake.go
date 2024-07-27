package packets

// HandShakePacket serverbound
type HandShakePacket struct {
	MinecraftPacket

	ProtocolVersion int32  `mc:"protocolVersion" json:"protocolVersion"` //VarInt
	ServerAddress   string `mc:"serverAddress" json:"serverAddress"`     //max: 255
	ServerPort      uint16 `mc:"serverPort" json:"serverPort"`
	NextState       int8   `mc:"nextState" json:"nextState"` //VarInt enum, 1 for status, 2 for login, 3 for transfer
}
