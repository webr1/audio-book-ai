package httpport

type HTTPServer interface {
	Use(middlewares ...Middleware)
	Run(address string) error
	Group(prefix string, middlewares ...Middleware) Group
	Init()
}
