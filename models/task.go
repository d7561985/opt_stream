package models

import "github.com/RichardKnop/uuid"

type Request struct {
	Data []Work `json:"data"`
}

type Work struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Temp float32   `json:"temp"`
	Hum  float32   `json:"hum"`
}
