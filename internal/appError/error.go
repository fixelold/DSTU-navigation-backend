package appError

type AppError struct {
	packageName string
	function    string
	message     string
	description string
	code        string
}

func NewError(packageName, function, message, description, code string) *AppError {
	return &AppError{
		packageName: packageName,
		function:    function,
		message:     message,
		description: description,
		code:        code,
	}
}

func (a *AppError) ChangeDescription(description string) *AppError {
	a.description = description
	return a
}

func (a *AppError) Error() string {
	return a.description
}
