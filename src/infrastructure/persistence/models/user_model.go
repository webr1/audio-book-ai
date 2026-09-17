package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model

	GoogleSub string `gorm:"column:google_sub;type:varchar(255);uniqueIndex;not null"`
	Email     string `gorm:"column:email;type:varchar(255);not null"`
	Name      string `gorm:"column:name;type:varchar(255);not null;default:''"`
	Picture   string `gorm:"column:picture;type:varchar(512);not null;default:''"`
}

func (UserModel) TableName() string {
	return "users"
}

func NewUserModel(googleSub, email, name, picture string) *UserModel {
	return &UserModel{GoogleSub: googleSub, Email: email, Name: name, Picture: picture}
}
