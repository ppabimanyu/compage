package exception

func (e *Exception) GetHttpCode() int {
	switch e.code {
	case InvalidParameterCode, InvalidDataCode:
		return 400
	case NotFoundCode:
		return 404
	case AlreadyExistsCode:
		return 409
	case PermissionDeniedCode:
		return 403
	case UnauthenticatedCode:
		return 401
	case InternalErrorCode:
		return 500
	default:
		return 500
	}
}
