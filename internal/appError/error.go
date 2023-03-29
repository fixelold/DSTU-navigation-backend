package appError

import "fmt"

type error interface {
	Error() string
}

type AppError struct {
	Context context
	Err error
}

type context struct {
	packageName string
	functionName string
}


func (a *AppError) Error() string {
	return fmt.Sprintf("Package: %s/n Function: %s/n Error: %v/n", a.Context.packageName, a.Context.functionName, a.Err)
}

func (a *AppError) WrapFunctionName(funcName string) {
	a.Context.functionName = a.Context.functionName + " -> " + funcName
}