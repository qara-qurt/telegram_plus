package model

import (
	"time"
)

type UserInfo struct {
	Username     string
	Login        string
	BirthDayDate time.Time
	Email        string
	Password     string
}

type User struct {
	ID           string
	Username     string
	Login        string
	BirthDayDate time.Time
	Email        string
	HashPassword string
	Status       *string
	Img          *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
