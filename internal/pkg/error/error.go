package error

type AppError struct {
	Code    int
	Message string
}

// ErrorNotFound creates a new AppError for "not found" errors.
func ErrorNotFound() *AppError {
	return &AppError{
		Code:    404,
		Message: "Resource not found",
	}
}

// InternalSystemError creates a new AppError for internal system errors.
func InternalSystemError() *AppError {
	return &AppError{
		Code:    500,
		Message: "Internal system error",
	}
}

// BadRequestError creates a new AppError for bad request errors.
func BadRequestError() *AppError {
	return &AppError{
		Code:    400,
		Message: "Bad request",
	}
}

// ErrorLimitExceeded creates a new AppError for limit exceeded errors.
func ErrorLimitExceeded() *AppError {
	return &AppError{
		Code:    429,
		Message: "Error limit exceeded",
	}
}
