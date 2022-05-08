package errors

type RequestValidationError struct {
	Message    map[string][]string `json:"message"`
	StatusCode int                 `json:"status_code"`
}

func NewRequestValidationError(messages map[string][]string, statusCode int) RequestValidationError {
	return RequestValidationError{
		Message:    messages,
		StatusCode: statusCode,
	}
}
