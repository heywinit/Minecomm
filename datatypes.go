package minecomm

//This file handles conversion of minecraft protocol datatypes to go incase some differ

// VarInt is a variable-length integer encoding used by Minecraft
type VarInt int32
type Boolean bool
type UnsignedByte uint8
type Short int16
type UnsignedShort uint16
type Int int32
type Long int64
type Float float32
type Double float64
type String string
type TextComponent string
type JSONTextComponent string
type Identifier string //max length: 32767