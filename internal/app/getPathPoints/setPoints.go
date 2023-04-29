package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

var (
	switchAxisError = appError.NewAppError("switch error")
	switchSignError = appError.NewAppError("switch error")
)

// точки от начала пути до вхождение в пределы сектора
func (d *data) setPointsPath2Sector(borderPoints, points, lastPathPoint models.Coordinates, axis int) (models.Coordinates) {
	p := models.Coordinates{
		X: (points.X),
		Y: (points.Y)}
	p.Widht = points.Widht
	p.Height = points.Height
	return p
}
