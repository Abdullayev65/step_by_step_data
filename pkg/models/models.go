package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Username      string `bun:",unique,nullzero,notnull" json:"username"`
	Password      string `bun:",nullzero,notnull" json:"password"`
	Admin         bool   `json:"-"`
}
type Article struct {
	bun.BaseModel `bun:"table:article"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Data          string `bun:",nullzero,notnull,type:text" json:"data"`
	UserID        int    `bun:",nullzero,notnull" json:"userID"`
	Step          int8   `bun:",nullzero,notnull" json:"step"`
	Active        bool   `bun:",notnull" json:"active"`
}
