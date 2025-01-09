package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(20)" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100)" validate:"required"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(72)" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(15)" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text" `
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(100)" validate:"required"`
	CreatedAt   time.Time `json:"-" gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `json:"-" gorm:"column:updated_at;type:timestamp"`
}

func (*User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type UserSession struct {
	ID                  int `gorm:"primary_key"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:varchar(255)" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(255)" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}
func (u *UserSession) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
