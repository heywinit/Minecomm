package types

import (
	"bytes"
	"github.com/heywinit/minecomm/internal/datatypes"
	"reflect"
)

func SerializeVarInt(field reflect.Value, databuf *bytes.Buffer) error {
	value, ok := field.Interface().(int32)
	if !ok {
		return ErrIncorrectFieldType
	}

	encodedData, _ := datatypes.EncodeVarInt(value)
	databuf.Write(encodedData)

	return nil
}

func SerializeVarLong(field reflect.Value, databuf *bytes.Buffer) error {
	value, ok := field.Interface().(int64)
	if !ok {
		return ErrIncorrectFieldType
	}

	encodedData, _ := datatypes.EncodeVarLong(value)
	databuf.Write(encodedData)

	return nil
}

func DeserializeVarInt(field reflect.Value, databuf *bytes.Buffer) error {
	decodedVal, _, err := datatypes.DecodeVarInt(databuf)
	if err != nil {
		return err
	}

	field.SetInt(int64(decodedVal))

	return nil
}

func DeserializeVarLong(field reflect.Value, databuf *bytes.Buffer) error {
	decodedVal, _, err := datatypes.DecodeVarLong(databuf)
	if err != nil {
		return err
	}

	field.SetInt(decodedVal)

	return nil
}
