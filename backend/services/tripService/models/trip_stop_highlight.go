package models

import "time"

type StopHighlight struct {
	Id          uint       `json:"id" example:"1"`
	CreatedAt   *time.Time `json:"created_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	UpdatedAt   *time.Time `json:"updated_at_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	Name        string     `json:"name" example:"Brandenburger Tor"`
	Description string     `json:"description" example:"The landmark of Berlin"`
	Long        float64    `json:"longitude" example:"52.520008"`
	Lat         float64    `json:"latitude" example:"13.404954"`
	TripStopID  uint       `json:"stop_id"`
}

type StopHighlightInput struct {
	Name        string  `json:"name" validate:"required" example:"Brandenburger Tor"`
	Description string  `json:"description" validate:"required" example:"The landmark of Berlin"`
	Long        float64 `json:"longitude" example:"52.520008"`
	Lat         float64 `json:"latitude" example:"13.404954"`
}
