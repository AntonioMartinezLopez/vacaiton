package models

import "time"

type TripQuery struct {
	Id              uint       `json:"id" example:"1"`
	CreatedAt       *time.Time `json:"created_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	UpdatedAt       *time.Time `json:"updated_at_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	Location        string     `json:"location" example:"Germany"`
	Duration        int        `json:"duration" example:"10"`
	Secrets         bool       `json:"secrets" gorm:"default:false"`
	MaximumDistance int        `json:"maximum_distance" example:"1000"`
	Focus           TripFocus  `json:"focus" sql:"type:enum('Cities','Nature','Mixed');default:'Mixed'" swaggertype:"string" enums:"Cities,Nature,Mixed" example:"Mixed"`
	TripID          uint       `json:"-"`
}

type CreateTripQueryInput struct {
	Location        string    `json:"location" binding:"required" example:"Germany" validate:"required"`
	Duration        int       `json:"duration" binding:"required" example:"10" validate:"required"`
	Secrets         bool      `json:"secrets" binding:"required" example:"true"`
	MaximumDistance int       `json:"maximum_distance" binding:"required" example:"100" validate:"required"`
	Focus           TripFocus `json:"focus" binding:"required" example:"Mixed" swaggertype:"string" enums:"Cities,Nature,Mixed"`
}

type UpdateTripQueryInput struct {
	Id              uint      `json:"id" example:"1" validate:"required"`
	Location        string    `json:"location" example:"Germany" validate:"required"`
	Duration        int       `json:"duration" example:"10" validate:"required"`
	Secrets         bool      `json:"secrets" example:"true" validate:"required"`
	MaximumDistance int       `json:"maximum_distance" example:"1000" validate:"required"`
	Focus           TripFocus `json:"focus" binding:"required" example:"Mixed" swaggertype:"string" enums:"Cities,Nature,Mixed"`
}
