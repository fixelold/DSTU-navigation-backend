package audToAud

import "navigation/internal/appError"

func (a *audToAudController) middle() appError.AppError {
	err := a.middletBuilding(a.points[1])
	if err.Err != nil {
		err.Wrap("middlePoints")
		return nil, err
	}
}