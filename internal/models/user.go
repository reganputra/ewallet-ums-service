package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id          int       `json:"id"`
	Username    string    `json:"username" gorm:"colum:username;type:varchar(20)" validate:"required"`
	Email       string    `json:"email" gorm:"colum:email;type:varchar(100)" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"colum:phone_number;type:varchar(15)" validate:"required"`
	FullName    string    `json:"full_name" gorm:"colum:full_name;type:varchar(100)" validate:"required"`
	Address     string    `json:"address" gorm:"colum:address;type:text" `
	Dob         string    `json:"dob" gorm:"colum:dob;type:date" `
	Password    string    `json:"password,omitempty" gorm:"colum:password;type:varchar(255)" validate:"required"`
	CreatedAt   time.Time `json:"-" `
	UpdatedAt   time.Time `json:"-" `
}

func (*User) TableName() string {
	return "users"
}
func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	Id                  int `gorm:"primaryKey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserId              uint      `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:varchar(500)" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(500)" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}
func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
