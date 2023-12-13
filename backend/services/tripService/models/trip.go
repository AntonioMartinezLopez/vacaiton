package models

import (
	"time"
)

type Trip struct {
	Id        uint       `json:"id" example:"1"`
	CreatedAt *time.Time `json:"created_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	Query     TripQuery  `json:"query" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Stops     []TripStop `json:"stops" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// type User struct {
// 	Id        int        `json:"id" example:"1"`
// 	Email     string     `json:"email" binding:"required" example:"testuser"`
// 	Firstname string     `json:"firstname" binding:"required" example:"testuser"`
// 	Lastname  string     `json:"lastname" binding:"required" example:"testuser"`
// 	Password  string     `json:"-"`
// 	CreatedAt *time.Time `json:"created_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
// 	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
// }

// type RegisterUserInput struct {
// 	Firstname string `json:"firstname" binding:"required" example:"testuser"`
// 	Lastname  string `json:"lastname" binding:"required" example:"testuser"`
// 	Email     string `json:"email" binding:"required" example:"testuser@test.de"`
// 	Password  string `json:"password" binding:"required" example:"test"`
// }

// type SignInUserInput struct {
// 	Email    string `json:"email" binding:"required" example:"testuser@test.de"`
// 	Password string `json:"password" binding:"required" example:"test"`
// }

// type RegisterUserOutput struct {
// 	UserId int `json:"userId" binding:"required" example:"1"`
// }
