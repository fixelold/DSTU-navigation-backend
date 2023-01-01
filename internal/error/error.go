package error

type Error struct {
	packageName string
	function    string
	description string
	error       error
	code        string
}

func NewError(packageName, function, description, code string, error error) *Error {
	return &Error{
		packageName: packageName,
		function:    function,
		description: description,
		error:       error,
		code:        code,
	}
}
