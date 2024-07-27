package packets

import "github.com/google/uuid"

type LoginStartPacket struct {
	MinecraftPacket

	Name string    `mc:"name" json:"name"` //max: 16
	UUID uuid.UUID `mc:"uuid" json:"uuid"`
}
