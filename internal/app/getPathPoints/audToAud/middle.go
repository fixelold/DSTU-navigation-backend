package audToAud

import "navigation/internal/appError"

func (a *audToAudController) middle() appError.AppError {
	err := a.middleBuilding()
	if err.Err != nil {
		err.Wrap("middlePoints")
		return err
	}

	return appError.AppError{}
}