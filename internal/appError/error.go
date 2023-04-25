package appError

import "fmt"
  
type AppError struct {
	Context     string
	Description string
	Err         error
}

func NewAppError(description string) *AppError {
	return &AppError{
		Context: "",
		Description: description,
	}
}

func (a *AppError) Error() string {
	return fmt.Sprintf("\n Context: %s\n Description: %s\n Error: %v\n", a.Context, a.Description, a.Err)
}

func (a *AppError) ToString() string {
	return fmt.Sprintf("%s", a.Err)
}

func (a *AppError) Wrap(funcName string) {
	if a.Context == "" {
		a.Context = funcName
	} else {
		a.Context = funcName + " -> " + a.Context	
	}
}
