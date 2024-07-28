package util

import "github.com/google/uuid"

func UUIDToBytes(uuid uuid.UUID) []byte {
	return uuid[:]
}
