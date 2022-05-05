package utils

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func NewErrJson(message string) *CustomError {
	return &CustomError{
		message: message,
	}
}
