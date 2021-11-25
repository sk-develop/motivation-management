package errors

type ValidationError struct {
	Message string
}

func NewValidationError(message string) *ValidationError {
	err := &ValidationError{Message: message}
	return err
}

func (e *ValidationError) Error() string {
	return e.Message
}
