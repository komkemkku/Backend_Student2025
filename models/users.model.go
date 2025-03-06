package model

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:users"`

	ID               int    `bun:",type:serial,autoincrement,pk"`
	Firstname        string `bun:"firstname"`
	Lastname         string `bun:"lastname"`
	Nickname         string `bun:"nickname"`
	Email            string `bun:"email"`
	Password         string `bun:"password"`
	StudentID        string `bun:"student_id"`
	Faculty          string `bun:"faculty"`
	MedicalCondition string `bun:"medical_condition"`
	FoodAllergies    string `bun:"food_allergies"`

	CreateUnixTimestamp
}
