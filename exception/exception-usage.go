package exception

// InvalidParameter creates a new Exception with the InvalidParameterCode error code.
// To be used when a parameter is invalid or does not meet the required criteria.
// The `err` parameter can be a map of field names to error messages for more detailed validation errors.
func InvalidParameter(message string, err map[string]string) *Exception {
	return _createException(InvalidParameterCode, message, nil, err)
}

// InvalidData creates a new Exception with the InvalidDataCode error code.
// To be used when the provided data is not valid.
func InvalidData(message string, err error) *Exception {
	return _createException(InvalidDataCode, message, err, nil)
}

// NotFound creates a new Exception with the NotFoundCode error code.
func NotFound(message string, err error) *Exception {
	return _createException(NotFoundCode, message, err, nil)
}

// AlreadyExists creates a new Exception with the AlreadyExistsCode error code.
func AlreadyExists(message string, err error) *Exception {
	return _createException(AlreadyExistsCode, message, err, nil)
}

// PermissionDenied creates a new Exception with the PermissionDeniedCode error code.
func PermissionDenied(message string, err error) *Exception {
	return _createException(PermissionDeniedCode, message, err, nil)
}

// Unauthenticated creates a new Exception with the UnauthenticatedCode error code.
func Unauthenticated(message string, err error) *Exception {
	return _createException(UnauthenticatedCode, message, err, nil)
}

// Internal creates a new Exception with the ErrorInternalCode error code.
// The original error that caused the exception is also included.
func Internal(message string, err error) *Exception {
	return _createException(InternalErrorCode, message, err, nil)
}
