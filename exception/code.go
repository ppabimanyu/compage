package exception

type Code string

const (
	InvalidParameterCode Code = "INVALID_PARAMETER"
	InvalidDataCode      Code = "INVALID_DATA"
	NotFoundCode         Code = "NOT_FOUND"
	AlreadyExistsCode    Code = "ALREADY_EXISTS"
	PermissionDeniedCode Code = "PERMISSION_DENIED"
	UnauthenticatedCode  Code = "UNAUTHENTICATED"
	InternalErrorCode    Code = "INTERNAL_ERROR"
)

func (e Code) ToString() string {
	return string(e)
}
