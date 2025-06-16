package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppabimanyu/compage/exception"
	"strings"
)

// Handler provides utility methods for handling JSON and XML responses in a Gin application.
// It includes methods for sending success responses, error responses, and handling exceptions.
type Handler struct {
}

// JSON sends a JSON response with the given response object.
// It also sets the RequestID in the response.
//
// Parameters:
// - c: The Fiber context.
// - r: The response object to send.
func (h *Handler) JSON(c *fiber.Ctx, r *Response) error {
	r.RequestID = GetCtxValueStr(c.Context(), RequestIDCtxKey)
	return c.Status(r.StatusCode).JSON(r)
}

// XML sends an XML response with the given response object.
// It also sets the RequestID in the response.
//
// Parameters:
// - c: The Fiber context.
// - r: The response object to send.
func (h *Handler) XML(c *fiber.Ctx, r *Response) error {
	r.RequestID = GetCtxValueStr(c.Context(), RequestIDCtxKey)
	return c.Status(r.StatusCode).XML(r)
}

func (h *Handler) ReturnByAccept(c *fiber.Ctx, r *Response) error {
	switch c.Get("Accept") {
	case "application/json":
		return h.JSON(c, r)
	case "application/xml":
		return h.XML(c, r)
	default:
		return h.JSON(c, r)
	}
}

func _successResponse() *Response {
	return &Response{
		StatusCode: 200,
		Message:    "success",
	}
}

// SuccessJSON sends a JSON response with a 200 status code and a success message.
//
// Parameters:
// - c: The Fiber context.
func (h *Handler) SuccessJSON(c *fiber.Ctx) error {
	r := _successResponse()
	return h.JSON(c, r)
}

// SuccessXML sends an XML response with a 200 status code and a success message.
//
// Parameters:
// - c: The Fiber context.
func (h *Handler) SuccessXML(c *fiber.Ctx) error {
	r := _successResponse()
	return h.XML(c, r)
}

func (h *Handler) Success(c *fiber.Ctx) error {
	r := _successResponse()
	return h.ReturnByAccept(c, r)
}

func _createdResponse() *Response {
	return &Response{
		StatusCode: 200,
		Message:    "created",
	}
}

// CreatedJSON sends a JSON response with a 201 status code and a success message.
//
// Parameters:
// - c: The Fiber context.
func (h *Handler) CreatedJSON(c *fiber.Ctx) error {
	r := _createdResponse()
	return h.JSON(c, r)
}

// CreatedXML sends an XML response with a 201 status code and a success message.
//
// Parameters:
// - c: The Fiber context.
func (h *Handler) CreatedXML(c *fiber.Ctx) error {
	r := _createdResponse()
	return h.XML(c, r)
}

func (h *Handler) Created(c *fiber.Ctx) error {
	r := _createdResponse()
	return h.ReturnByAccept(c, r)
}

func _dataResponse(data any) *Response {
	return &Response{
		StatusCode: 200,
		Message:    "success",
		Data:       data,
	}
}

// DataJSON sends a JSON response with a 200 status code, a success message, and the provided data.
//
// Parameters:
// - c: The Fiber context.
// - data: The data to include in the response.
func (h *Handler) DataJSON(c *fiber.Ctx, data any) error {
	r := _dataResponse(data)
	return h.JSON(c, r)
}

// DataXML sends an XML response with a 200 status code, a success message, and the provided data.
//
// Parameters:
// - c: The Fiber context.
// - data: The data to include in the response.
func (h *Handler) DataXML(c *fiber.Ctx, data any) error {
	r := _dataResponse(data)
	return h.XML(c, r)
}

func (h *Handler) Data(c *fiber.Ctx, data any) error {
	r := _dataResponse(data)
	return h.ReturnByAccept(c, r)
}

func _exceptionResponse(exc *exception.Exception) *Response {
	var detailErr any
	if exc.GetCode() == exception.InvalidParameterCode {
		detailErr = exc.GetErrorMap()
	} else {
		detailErr = exc.GetError()
	}
	r := &Response{
		StatusCode: exc.GetHttpCode(),
		Message:    exc.GetMessage(),
		Error: &Error{
			Code:    exc.GetCode().ToString(),
			Details: detailErr,
		},
	}
	return r
}

// ExceptionJSON sends a JSON response for the given exception and aborts the request.
//
// Parameters:
// - c: The Fiber context.
// - exc: The exception to handle.
func (h *Handler) ExceptionJSON(c *fiber.Ctx, exc *exception.Exception) error {
	r := _exceptionResponse(exc)
	return h.JSON(c, r)
}

// ExceptionXML sends an XML response for the given exception and aborts the request.
//
// Parameters:
// - c: The Fiber context.
// - exc: The exception to handle.
func (h *Handler) ExceptionXML(c *fiber.Ctx, exc *exception.Exception) error {
	r := _exceptionResponse(exc)
	return h.XML(c, r)
}

func (h *Handler) Exception(c *fiber.Ctx, exc *exception.Exception) error {
	r := _exceptionResponse(exc)
	return h.ReturnByAccept(c, r)
}

func _badRequestResponse(msg string, err error) *Response {
	r := &Response{
		StatusCode: 400,
		Message:    msg,
	}
	if err != nil {
		r.Error = &Error{
			Code:    "BAD_REQUEST",
			Details: err.Error(),
		}
	}
	return r
}

// BadRequestJSON sends a JSON response with a 400 status code and the provided message and error details.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) BadRequestJSON(c *fiber.Ctx, msg string, err error) error {
	r := _badRequestResponse(msg, err)
	return h.JSON(c, r)
}

// BadRequestXML sends an XML response with a 400 status code and the provided message and error details.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) BadRequestXML(c *fiber.Ctx, msg string, err error) error {
	r := _badRequestResponse(msg, err)
	return h.XML(c, r)
}

func (h *Handler) BadRequest(c *fiber.Ctx, msg string, err error) error {
	r := _badRequestResponse(msg, err)
	return h.ReturnByAccept(c, r)
}

// NotFoundJSON sends a JSON response for a not found error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) NotFoundJSON(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionJSON(c, exception.NotFound(msg, err))
}

// NotFoundXML sends an XML response for a not found error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) NotFoundXML(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionXML(c, exception.NotFound(msg, err))
}

func (h *Handler) NotFound(c *fiber.Ctx, msg string, err error) error {
	return h.Exception(c, exception.NotFound(msg, err))
}

// ForbiddenJSON sends a JSON response for a forbidden error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) ForbiddenJSON(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionJSON(c, exception.PermissionDenied(msg, err))
}

// ForbiddenXML sends an XML response for a forbidden error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) ForbiddenXML(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionXML(c, exception.PermissionDenied(msg, err))
}

func (h *Handler) Forbidden(c *fiber.Ctx, msg string, err error) error {
	return h.Exception(c, exception.PermissionDenied(msg, err))
}

// UnauthorizedJSON sends a JSON response for an unauthorized error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) UnauthorizedJSON(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionJSON(c, exception.Unauthenticated(msg, err))
}

// UnauthorizedXML sends an XML response for an unauthorized error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) UnauthorizedXML(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionXML(c, exception.Unauthenticated(msg, err))
}

func (h *Handler) Unauthorized(c *fiber.Ctx, msg string, err error) error {
	return h.Exception(c, exception.Unauthenticated(msg, err))
}

// InternalServerErrorJSON sends a JSON response for an internal server error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) InternalServerErrorJSON(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionJSON(c, exception.Internal(msg, err))
}

// InternalServerErrorXML sends an XML response for an internal server error.
//
// Parameters:
// - c: The Fiber context.
// - msg: The error message.
// - err: Additional error details.
func (h *Handler) InternalServerErrorXML(c *fiber.Ctx, msg string, err error) error {
	return h.ExceptionXML(c, exception.Internal(msg, err))
}

func (h *Handler) InternalServerError(c *fiber.Ctx, msg string, err error) error {
	return h.Exception(c, exception.Internal(msg, err))
}

func (h *Handler) GetTokenFromHeader(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if len(authHeader) > 7 && strings.HasPrefix(authHeader, "Bearer ") {
		return authHeader[7:]
	}
	return ""
}
