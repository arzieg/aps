package aps

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID       uuid.UUID
	Name     string
	Location string
	Price    float64
	Date     time.Time
}
