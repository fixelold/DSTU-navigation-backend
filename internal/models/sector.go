package models

import "time"

type Sector struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Number uint `json:"number"`
	FloorID int `json:"id_floor"`
}
