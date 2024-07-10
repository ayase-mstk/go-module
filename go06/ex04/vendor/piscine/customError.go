package piscine

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func (e *CustomError) setMsg(msg string) {
	e.Message = msg
}
