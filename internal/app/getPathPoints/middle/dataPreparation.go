package middle

import (
	"navigation/internal/models"
)

func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	var path models.Coordinates

	if axis == m.constData.axisX {
		if points.Height == -10 {path = m.preparationDownX(borderPoint, points)
		}else if points.Height == 10 {path = m.preparationUpX(borderPoint, points)}
	
	}else if axis == m.constData.axisY {
		if points.Widht == -10 {path = m.preparationRightY(borderPoint, points)
			}else if points.Widht == 10 {path = m.preparationLeftY(borderPoint, points)}
	}

	return path
}

func (m *middleController) finalPreparation(axis int, borderPoint, points models.Coordinates, exceptions bool) models.Coordinates {
	var path models.Coordinates
	if axis == m.constData.axisX {
		if exceptions {path = m.leftAndRightX(borderPoint, points)
		} else if points.Height == -10 {path = m.downX(borderPoint, points)
		} else if points.Height == 10 {path = m.upX(borderPoint, points)}

	} else if axis == m.constData.axisY {
		if exceptions {path = m.upAndDownY(borderPoint, points)
			} else if points.Widht == -10 {path = m.rightY(borderPoint, points)
			} else if points.Widht == 10 {path = m.leftY(borderPoint, points)}
	}

	return path
}