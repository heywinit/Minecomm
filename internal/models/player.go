package models

import "github.com/google/uuid"

type Player struct {
	Name string    `mc:"name" json:"name"`
	UUID uuid.UUID `mc:"uuid" json:"uuid"`
}
