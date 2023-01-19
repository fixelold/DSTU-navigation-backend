package models

import "time"

type SectorLink struct {
	ID           uint `json:"id"`
	CreatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	NumberSector int       `json:"number_sector"`
	NumberLink   int       `json:"link"`
}
