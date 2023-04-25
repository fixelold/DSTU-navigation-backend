package models

type ImportantPlaces struct {
	ID int `json:"id"`
	Name string `json:"name" binding:"required"`
	AuditoryID int `json:"auditory_id" binding:"required"`
}