package ctx

import (
	"strconv"

	"audio-book-ai/src/core/application/response"
)

// GetBody binds and validates the request body into T, panicking with a
// *response.Response on failure so the recovery middleware turns it into
// a proper JSON error response.
func GetBody[T any](c Context) *T {
	var body T
	if err := c.Bind(&body); err != nil {
		panic(response.NewFailResponse(400, "invalid request body"))
	}
	if err := c.Validate(&body); err != nil {
		panic(response.NewFailResponse(400, err.Error()))
	}
	return &body
}

func GetUintPathParam(c Context, name string) uint {
	v, err := strconv.ParseUint(c.Param(name), 10, 64)
	if err != nil {
		panic(response.NewFailResponse(400, "invalid path param: "+name))
	}
	return uint(v)
}
