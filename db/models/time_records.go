package models

import (
	"time"
)

// Base contains common columns for all tables.
type TimeRecords struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
