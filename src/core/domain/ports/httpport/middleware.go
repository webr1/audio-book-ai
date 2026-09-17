package httpport

type Middleware func(next Handler) Handler
