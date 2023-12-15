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
	UserId    string     `json:"user_id"`
}
