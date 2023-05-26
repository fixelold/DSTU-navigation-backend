package audToAud

import (
	"fmt"

	"navigation/internal/models"
)

// func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
// 	var path models.Coordinates
// 	if axis == m.constData.axisX {
// 		if points.Widht == -5 {path = m.preparationDownX(borderPoint, points)
// 		}else if points.Widht == 5 {path = m.preparationUpX(borderPoint, points)}

// 		if points.Widht == -10 {path = m.preparationDownX(borderPoint, points)
// 			}else if points.Widht == 10 {path = m.preparationUpX(borderPoint, points)}

// 	}else if axis == m.constData.axisY {
// 		if points.Widht == -5 {path = m.preparationRightY(borderPoint, points)
// 			}else if points.Widht == 5 {path = m.preparationLeftY(borderPoint, points)}

// 		if points.Widht == -10 {path = m.preparationRightY(borderPoint, points)
// 			}else if points.Widht == 10 {path = m.preparationLeftY(borderPoint, points)}
// 	}

// 	return path
// }

func (m *middleController) finalPreparation(axis int, borderPoint, points models.Coordinates, exceptions bool) models.Coordinates {
	var path models.Coordinates
	var factor int
	fmt.Println("work: ", borderPoint, points, axis)
	if exceptions {
		if axis == m.constData.axisX {
			// Проверка сверху начальная аудитория или снизу
			if points.Height == -10  {factor = -1 // снизу
			} else if points.Height == 10  {factor = 1} else {factor = 1} // сверху

			if borderPoint.X > points.X {path = m.pointsRightX(borderPoint, points, m.endPoints,  factor)
				} else if borderPoint.X < points.X {path = m.pointsLeftX(borderPoint, points, m.endPoints, factor)}

		} else if axis == m.constData.axisY {
			if points.Widht == -10  {factor = -1
				} else if points.Widht == 10  {factor = 1}

			if borderPoint.Y > points.Y {path = m.pointsDownY(borderPoint, points, m.endPoints, factor)
				} else if borderPoint.Y < points.Y {path = m.pointsUpY(borderPoint, points, m.endPoints, factor)}
		}

	} else {
		if axis == m.constData.axisX {path = m.firstFinalX(borderPoint, points, factor)
		} else if axis == m.constData.axisY {path = m.firstFinalY(borderPoint, points, factor)}
	}

	return path
}
