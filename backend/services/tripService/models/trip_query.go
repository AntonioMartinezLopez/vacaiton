package models

import "time"

type TripQuery struct {
	Id              uint       `json:"id" example:"1"`
	CreatedAt       *time.Time `json:"created_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	UpdatedAt       *time.Time `json:"updated_at_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	Country         string     `json:"country" example:"Germany"`
	Duration        int        `json:"duration" example:"10"`
	Secrets         bool       `json:"secrets" gorm:"default:false"`
	MaximumDistance int        `json:"maximum_distance" example:"1000"`
	Focus           TripFocus  `json:"focus" sql:"type:enum('cities','nature','mixed');default:'mixed'" gorm:"embedded"`
	TripID          uint       `json:"-"`
}

type CreateTripQueryInput struct {
	Country         string    `json:"country" binding:"required" example:"Germany" validate:"required"`
	Duration        int       `json:"duration" binding:"required" example:"10" validate:"required"`
	Secrets         bool      `json:"secrets" binding:"required" validate:"required"`
	MaximumDistance int       `json:"maximum_distance" binding:"required" example:"1000" validate:"required"`
	Focus           TripFocus `json:"focus" binding:"required"`
}

type UpdateTripQueryInput struct {
	Id              uint      `json:"id" example:"1" validate:"required"`
	Country         string    `json:"country" example:"Germany" validate:"required"`
	Duration        int       `json:"duration" example:"10" validate:"required"`
	Secrets         bool      `json:"secrets" validate:"required"`
	MaximumDistance int       `json:"maximum_distance" example:"1000" validate:"required"`
	Focus           TripFocus `json:"focus"`
}
