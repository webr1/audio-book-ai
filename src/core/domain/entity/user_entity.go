package entity

import "time"

type UserEntity struct {
	ID        uint
	GoogleSub string
	Email     string
	Name      string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserEntity(googleSub, email, name, picture string) *UserEntity {
	return &UserEntity{GoogleSub: googleSub, Email: email, Name: name, Picture: picture}
}
