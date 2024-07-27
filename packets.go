package minecomm

import (
	"github.com/heywinit/minecomm/internal/models/packets"
	"io"
)

// SerializablePacket defines the standard methods that a struct should have
// in order to be serializable by the library
//
// You can actually create your own methods as long as they respect this standard
type SerializablePacket interface {
	// SerializeData takes an interface pointer as input and serializes all the fields in the
	// data buffer. It can and will return an error in case of invalid data
	SerializeData(inter interface{}) error

	SerializeCompressed(writer io.Writer, compressionThreshold int) error
	SerializeUncompressed(writer io.Writer) error
}

// WritePacket calls SerializeData and then calls WriteRawPacket
func (mc *Client) WritePacket(packet SerializablePacket) error {
	mc.writeChannel.Lock()
	defer mc.writeChannel.Unlock()

	if err := packet.SerializeData(packet); err != nil {
		return err
	}

	if mc.IsCompressionEnabled() {
		return packet.SerializeCompressed(mc.con, int(mc.compressionThreshold))
	} else {
		return packet.SerializeUncompressed(mc.con)
	}
}

// WriteRawPacket takes a rawpacket as input and serializes it in the con
func (mc *Client) WriteRawPacket(rawPacket *packets.MinecraftRawPacket) error {
	mc.writeChannel.Lock()
	defer mc.writeChannel.Unlock()

	if mc.IsCompressionEnabled() {
		return rawPacket.WriteCompressed(mc.con)
	} else {
		return rawPacket.WriteUncompressed(mc.con)
	}
}

// ReceiveRawPacket reads a raw packet from the con but doesn't deserialize
// neither uncompress it
func (mc *Client) ReceiveRawPacket() (*packets.MinecraftRawPacket, error) {
	mc.readChannel.Lock()
	defer mc.readChannel.Unlock()

	if mc.IsCompressionEnabled() {
		return packets.FromCompressedReader(mc.con)
	} else {
		return packets.FromUncompressedReader(mc.con)
	}
}

// ReceivePacket receives and deserializes a packet from the con, uncompressing it
// if necessary
func (mc *Client) ReceivePacket() (*packets.MinecraftPacket, error) {
	rawPacket, err := mc.ReceiveRawPacket()
	if err != nil {
		return nil, err
	}

	return packets.FromRawPacket(rawPacket)
}
