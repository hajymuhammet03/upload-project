package language

import "time"

type UUID struct {
	UUID string `json:"uuid"`
}

type Language struct {
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	LastUpdate time.Time `json:"last_update"`
}
