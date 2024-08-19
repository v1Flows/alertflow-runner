package models

import (
	"time"

	"github.com/google/uuid"
)

type ExecutionSteps struct {
	ID             uuid.UUID `json:"id"`
	ExecutionID    string    `json:"execution_id"`
	ActionName     string    `json:"action_name"`
	ActionID       string    `json:"action_id"`
	ActionMessages []string  `json:"action_messages"`
	Error          bool      `json:"error"`
	PatternMatch   bool      `json:"pattern_match"`
	Finished       bool      `json:"finished"`
	Paused         bool      `json:"paused"`
	StartedAt      time.Time `json:"started_at"`
	FinishedAt     time.Time `json:"finished_at"`
	NoResult       bool      `json:"no_result"`
}
