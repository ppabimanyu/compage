package exception

func (e *Exception) GetGRPCCode() int32 {
	switch e.code {
	case InvalidParameterCode, InvalidDataCode:
		return 3
	case NotFoundCode:
		return 5
	case AlreadyExistsCode:
		return 6
	case PermissionDeniedCode:
		return 7
	case UnauthenticatedCode:
		return 16
	case InternalErrorCode:
		return 13
	default:
		return 13
	}
}
