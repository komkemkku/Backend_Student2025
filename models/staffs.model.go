package model

import "github.com/uptrace/bun"

type Staffs struct {
	bun.BaseModel `bun:"table:staffs"`

	ID       int    `bun:",type:serial,autoincrement,pk"`
	Username string `bun:"username"`
	Password string `bun:"password"`

	CreateUnixTimestamp
}
