package models

import "time"

type BorderPoint struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	X         int        `json:"x"`
	Y         int        `json:"y"`
	Widht     int        `json:"widht"`
	Height    int        `json:"height"`
}
