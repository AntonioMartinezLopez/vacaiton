package models

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Email     string     `json:"email" binding:"required"`
	Firstname string     `json:"firstname" binding:"required"`
	Lastname  string     `json:"lastname" binding:"required"`
	Password  string     `json:"-"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}

type RegisterUserInput struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type SignUpUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
