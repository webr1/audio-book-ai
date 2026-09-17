package httpport

type Group interface {
	Use(middlewares ...Middleware)
	GET(path string, handler Handler, middlewares ...Middleware)
	POST(path string, handler Handler, middlewares ...Middleware)
	PUT(path string, handler Handler, middlewares ...Middleware)
	DELETE(path string, handler Handler, middlewares ...Middleware)
	PATCH(path string, handler Handler, middlewares ...Middleware)
	Any(path string, handler Handler, middlewares ...Middleware)
}
