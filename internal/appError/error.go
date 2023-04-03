package appError

import "fmt"

type AppError struct {
	Context string
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

func (a *AppError) Wrap(packageName, fileName, functionName string) {
	context := fmt.Sprintf("( %s -> %s -> %s)", packageName, fileName, functionName)
	if a.Context == "" {
		a.Context = context
	} else {
		a.Context = context + " -> " + a.Context	
	}
}
