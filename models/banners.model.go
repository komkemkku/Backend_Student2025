package model

import (
	"github.com/uptrace/bun"
)

type Banner struct {
	bun.BaseModel `bun:"table:banners"`

	ID        int    `bun:",type:serial,autoincrement,pk"`
	Banner    string `bun:"banner_image"`
	Published string `bun:"published"`

	CreateUnixTimestamp
}
