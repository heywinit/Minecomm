package datatypes

import (
	"bytes"
	"errors"
	"io"
)

var ErrVarLongTooLarge = errors.New("minecomm: varlong too large")

// EncodeVarLong takes an in64, and encodes it into an array of bytes of size n as specified on https://wiki.vg/Data_types#VarInt_and_VarLong
func EncodeVarLong(inputValue int64) (varint []byte, n int) {
	val := uint64(inputValue)

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

func DecodeVarLong(reader io.Reader) (result int64, numRead int, err error) {
	read := make([]byte, 1)

	for {
		_, err = io.ReadFull(reader, read)
		if err != nil {
			return
		}

		readByte := read[0]

		val := int64(readByte & 0b01111111)
		result |= val << (7 * numRead)

		numRead++
		if numRead > 10 {
			err = ErrVarLongTooLarge
			return
		}

		if (readByte & 0b10000000) == 0 {
			break
		}
	}

	return result, numRead, nil
}
