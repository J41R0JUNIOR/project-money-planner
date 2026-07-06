package errors

import "fmt"

type AppError struct {
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
