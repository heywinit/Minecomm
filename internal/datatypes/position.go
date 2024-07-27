package datatypes

import "encoding/binary"

// Position represents a 3D position in the Minecraft world
type Position struct {
	X int32
	Y int16
	Z int32
}

// Encode converts a Position to its 64-bit representation
func (p *Position) Encode() []byte {
	val := (int64(p.X)&0x3FFFFFF)<<38 | (int64(p.Z)&0x3FFFFFF)<<12 | int64(p.Y&0xFFF)

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(val))
	return buf
}

// Decode converts a 64-bit binary representation to a Position
func Decode(data []byte) *Position {
	if len(data) != 8 {
		return nil // or handle error appropriately
	}

	val := int64(binary.BigEndian.Uint64(data))

	x := int32(val >> 38)
	y := int16(val << 52 >> 52)
	z := int32(val << 26 >> 38)

	// Apply sign extension
	if x >= 1<<25 {
		x -= 1 << 26
	}
	if int32(y) >= 1<<11 {
		y -= 1 << 12
	}
	if z >= 1<<25 {
		z -= 1 << 26
	}

	return &Position{X: x, Y: y, Z: z}
}
