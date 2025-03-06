package model

import (
	"github.com/uptrace/bun"
)

type Tickets struct {
	bun.BaseModel `bun:"table:tickets"`

	ID      int    `bun:",type:serial,autoincrement,pk"`
	UserID  int    `bun:"user_id"`
	EventID int    `bun:"event_id"`
	QrCode  string `bun:"qr_code"`

	CreateUnixTimestamp
}
