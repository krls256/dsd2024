package entities

import "github.com/google/uuid"

type Log struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
