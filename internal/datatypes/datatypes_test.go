package datatypes_test

import (
	"bytes"
	"github.com/heywinit/minecomm/internal/datatypes"
	"testing"
)

func TestPosition(t *testing.T) {
	// Test Encode
	pos := &datatypes.Position{X: 1, Y: 2, Z: 3}
	encoded := pos.Encode()
	if len(encoded) != 8 {
		t.Errorf("Expected 8 bytes, got %d", len(encoded))
	}
	decoded := datatypes.Decode(encoded)
	if decoded.X != pos.X || decoded.Y != pos.Y || decoded.Z != pos.Z {
		t.Errorf("Expected %v, got %v", pos, decoded)
	}
}

func TestVarInt(t *testing.T) {
	// Test Encode
	varint, n := datatypes.EncodeVarInt(25565)
	if n != 3 {
		t.Errorf("Expected 3 bytes, got %d", n)
	}
	decoded, _, _ := datatypes.DecodeVarInt(bytes.NewReader(varint))
	if decoded != 25565 {
		t.Errorf("Expected 25565, got %d", decoded)
	}
}

func TestVarLong(t *testing.T) {
	// Test Encode
	varlong, n := datatypes.EncodeVarLong(2147483647)
	if n != 5 {
		t.Errorf("Expected 5 bytes, got %d", n)
	}
	decoded, _, _ := datatypes.DecodeVarLong(bytes.NewReader(varlong))
	if decoded != 2147483647 {
		t.Errorf("Expected 2147483647, got %d", decoded)
	}
}
