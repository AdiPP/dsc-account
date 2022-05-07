package errors

type ServiceError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewServiceError(message string, statusCode int) ServiceError {
	return ServiceError{
		Message:    message,
		StatusCode: statusCode,
	}
}
