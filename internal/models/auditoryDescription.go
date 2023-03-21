package models

type AuditoryDescription struct {
	ID uint
	AuditoryID uint `json:"auditory_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}