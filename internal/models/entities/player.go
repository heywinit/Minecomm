package entities

import "github.com/google/uuid"

type Player struct {
	Name string    `json:"name"`
	UUID uuid.UUID `json:"uuid"`
}
