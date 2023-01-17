package models

import "time"

type Auditorium struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Number string `json:"number"`
	SectorID int `json:"id_sector"`
}
