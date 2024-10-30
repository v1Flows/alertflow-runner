package models

import (
	"encoding/json"
	"time"
)

type Register struct {
	Registered                bool            `json:"registered"`
	AvailableActions          json.RawMessage `json:"available_actions"`
	AvailablePayloadInjectors json.RawMessage `json:"available_payload_injectors"`
	LastHeartbeat             time.Time       `json:"last_heartbeat"`
	Version                   string          `json:"version"`
	Mode                      string          `json:"mode"`
}
