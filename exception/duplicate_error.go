package exception

type DuplicateError struct {
	Message string
}

func (duplicateError DuplicateError) Error() string {
	return duplicateError.Message
}