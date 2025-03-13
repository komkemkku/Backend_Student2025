package model

import (
	"github.com/uptrace/bun"
)

type Events struct {
	bun.BaseModel `bun:"table:events"`

	ID          int    `bun:",type:serial,autoincrement,pk"`
	Image       string   `bun:"image"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
	Location    string `bun:"location"`
	Dress       string `bun:"dress"`
	StartTime   string `bun:"start_time"`
	EndTime     string `bun:"end_time"`
	StartDate   string `bun:"start_date"`
	EndDate     string `bun:"end_date"`
	Is_active   bool   `bun:"is_active"`

	CreateUnixTimestamp
}
