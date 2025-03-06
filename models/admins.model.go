package model

import "github.com/uptrace/bun"

type Admins struct {
	bun.BaseModel `bun:"table:admins"`

	ID       int    `bun:",type:serial,autoincrement,pk"`
	Username string `bun:"username"`
	Password string `bun:"password"`

	CreateUnixTimestamp
}
