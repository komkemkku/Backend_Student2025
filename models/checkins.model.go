package model

import "github.com/uptrace/bun"

type Checkins struct {
	bun.BaseModel `bun:"table:checkins"`

	ID       int `bun:",type:serial,autoincrement,pk"`
	TicketID int `bun:"ticket_id"`
	StaffID  int `bun:"staff_id"`

	CreateUnixTimestamp
}
