package audToTransition

import (
	"navigation/internal/models"
)

// для расчета пути, если он не входит в границы следующего сектора
func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	var path models.Coordinates
	if axis == m.constData.axisX {
		if points.Widht == -5 {path = m.preparationDownX(borderPoint, points)
		}else if points.Widht == 5 {path = m.preparationUpX(borderPoint, points)}
	
	}else if axis == m.constData.axisY {
		if points.Widht == -5 {path = m.preparationRightY(borderPoint, points) // поменял на 5
			}else if points.Widht == 5 {path = m.preparationLeftY(borderPoint, points)} // поменял на 5

		if points.Widht == -10 {path = m.preparationRightY(borderPoint, points)
			}else if points.Widht == 10 {path = m.preparationLeftY(borderPoint, points)}
	}

	return path
}

// для расчета пути, если он входит в границы следующего сектора
func (m *middleController) finalPreparation(axis int, borderPoint, points models.Coordinates, exceptions bool) (models.Coordinates, bool) {
	var path models.Coordinates
	var lenght = len(m.Points)

	if lenght == 1 {
		if axis == m.constData.axisX {
			if exceptions {path = m.leftAndRightX(borderPoint, points)
			} else if points.Height == -10 {path = m.downX(borderPoint, points)
			} else if points.Height == 10 {path = m.upX(borderPoint, points)}
	
		} else if axis == m.constData.axisY {
			if exceptions {path = m.upAndDownY(borderPoint, points)
				} else if points.Widht == -5 {path = m.rightY(borderPoint, points)
				} else if points.Widht == 5 {path = m.leftY(borderPoint, points)}
		}

	} else {
		if axis == m.constData.axisX {
			if points.Height < 0 {
				if points.Widht == 5 {path = m.downLeftX(borderPoint, points) // убрал m.Points[0] в самом начале
				} else if points.Widht == -5 {path = m.downRightX(borderPoint, points)} // убрал m.Points[0] в самом начале

			} else if points.Height > 0 {
				if points.Widht == 5 {path = m.upLeftX(borderPoint, points) // убрал m.Points[0] в самом начале
					} else if points.Widht == -5 {path = m.upRightX(borderPoint, points)} // убрал m.Points[0] в самом начале
			}
		} else if axis == m.constData.axisY {
			if points.Widht > 0 {
				if points.Height == 5 {path = m.leftDownY(borderPoint, points)
				} else if points.Height == -5 {path = m.leftUpY(borderPoint, points)}

			} else if points.Widht < 0 {
				if points.Height == 5 {path = m.rightDownY(borderPoint, points)
					} else if points.Height == -5 {path = m.rightUpY(borderPoint, points)}
			}
		}
	}	

	return path, true
}