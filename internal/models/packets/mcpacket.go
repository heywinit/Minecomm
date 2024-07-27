package packets

import (
	"bytes"
	"compress/zlib"
	"github.com/heywinit/minecomm/internal/datatypes"
	"io"
)

type MinecraftPacket struct {
	Length   int32  `mc:"length" json:"length"`     //VarInt, length of the packet
	PacketID int32  `mc:"packetID" json:"packetID"` //VarInt, different for each type of packet
	Data     []byte `mc:"data" json:"data"`
}

// SerializeUncompressed serializes the fields into a buffer and
// writes a complete packet in writer using the uncompressed format.
// Never uses compression
func (p *MinecraftPacket) SerializeUncompressed(writer io.Writer) error {
	packetId, _ := datatypes.EncodeVarInt(p.PacketID)

	packet := new(MinecraftRawPacket)
	packet.data = append(packetId, p.Data...)

	return packet.WriteUncompressed(writer)
}

// SerializeCompressed serializes the fields into a buffer, and if the buffer exceeds compressionTreshold
// in length, it proceeds to compress it using zlib. Writes a complete packet into writer
func (p *MinecraftPacket) SerializeCompressed(writer io.Writer, compressionTreshold int) error {
	encPid, pIdLen := datatypes.EncodeVarInt(p.PacketID)

	uncompressedBuffer := new(bytes.Buffer)
	uncompressedBuffer.Write(encPid)
	uncompressedBuffer.Write(p.Data)

	dataLength := pIdLen + len(p.Data)

	packet := new(MinecraftRawPacket)

	if dataLength >= compressionTreshold {
		// Create a new compressed buffer and a zlib writer pointing to it
		compressedData := new(bytes.Buffer)
		dataWriter := zlib.NewWriter(compressedData)

		// write into the zlib writer copying from the uncompressed buffer
		if _, err := io.Copy(dataWriter, uncompressedBuffer); err != nil {
			return err
		}

		// flush the zlib writer, filling up compressedData
		if err := dataWriter.Close(); err != nil {
			return err
		}

		packet.dataLength = dataLength
		packet.data = compressedData.Bytes()

		return packet.WriteCompressed(writer)
	} else {
		packet.dataLength = 0
		packet.data = uncompressedBuffer.Bytes()

		return packet.WriteCompressed(writer)
	}
}

// FromRawPacket returns a new MinecraftPacket, and takes a rawpacket as an input.
// rawPacket.ReadAll() is called to decode the packetid and the packet data
func FromRawPacket(rawPacket *MinecraftRawPacket) (*MinecraftPacket, error) {
	pId, data, err := rawPacket.ReadAll()
	if err != nil {
		return nil, err
	}

	packet := new(MinecraftPacket)
	packet.PacketID = pId
	packet.Data = data

	return packet, nil
}

func (l MinecraftPacket) SerializeData(_ interface{}) error {
	return nil
}
