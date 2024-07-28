package packets

type HandshakePacket struct {
	MinecraftPacket

	ProtocolVersion int32  `mc:"varint"`
	ServerAddress   string `mc:"string"`
	ServerPort      uint16 `mc:"inherit"`
	NextState       int32  `mc:"varint"`
}
