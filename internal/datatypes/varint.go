package datatypes

import (
	"bytes"
	"errors"
	"io"
)

var ErrVarIntTooLarge = errors.New("minecomm: varint too large")

// EncodeVarInt takes an int32 and encodes it into an array of bytes of n size as specified on https://wiki.vg/Data_types#VarInt_and_VarLong
func EncodeVarInt(input int32) (varint []byte, n int) {
	val := uint32(input)

	buffer := new(bytes.Buffer)
	for {
		temp := (byte)(val & 0b01111111)
		val >>= 7

		if val != 0 {
			temp |= 0b10000000
		}

		buffer.WriteByte(temp)
		n++

		if val == 0 {
			break
		}
	}

	return buffer.Bytes(), n
}

func DecodeVarInt(reader io.Reader) (result int32, numRead int, err error) {
	read := make([]byte, 1)

	for {
		_, err = io.ReadFull(reader, read)
		if err != nil {
			return
		}

		readByte := read[0]

		val := int32(readByte & 0b01111111)
		result |= val << (7 * numRead)

		numRead++
		if numRead > 5 {
			err = ErrVarIntTooLarge
			return
		}

		if (readByte & 0b10000000) == 0 {
			break
		}
	}

	return result, numRead, nil
}
