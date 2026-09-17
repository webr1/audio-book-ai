package response

type Response struct {
	Status  int    `json:"status,omitempty"`
	Code    Code   `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse(status int, code Code, message string, data any) *Response {
	return &Response{Status: status, Code: code, Message: message, Data: data}
}

func NewFailResponse(status int, message string) *Response {
	return &Response{Status: status, Message: message}
}

func NewOptionalResponse(status int, code Code, message string) *Response {
	return NewResponse(status, code, message, nil)
}

func (r *Response) Error() string {
	return r.Message
}

var (
	NotFoundError         = NewOptionalResponse(404, CodeNotFound, "not found")
	ConflictError         = NewOptionalResponse(409, CodeConflict, "conflict error")
	UnauthorizedError     = NewFailResponse(401, "unauthorized")
	InvalidTokenError     = NewFailResponse(401, "invalid token")
	ExpiredTokenError     = NewFailResponse(401, "token expired")
	PermissionDeniedError = NewFailResponse(403, "permission denied")
)
