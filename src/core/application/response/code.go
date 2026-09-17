package response

type Code int

const (
	CodeNone         Code = 0
	CodeNotFound     Code = 4104
	CodeConflict     Code = 4109
	CodeAlreadyExist Code = 4112
)
