package apperr

type Level int

const (
	Info Level = iota + 1
	Warn
	Fatal
)

type Error struct {
	error

	log   string
	level Level
	code  Code
}

func CreateError(code Code, level Level, log string) *Error {
	return &Error{
		code:  code,
		level: level,
		log:   log,
	}
}

func (e Error) Error() string {
	return string(e.code)
}

func (e Error) String() string {
	return e.Error()
}

func (e Error) Message() string {
	return e.code.message()
}

func (e Error) Level() Level {
	return e.level
}

func (e Error) Log() string {
	return e.log
}

func (e Error) Status() int {
	return e.code.status()
}

func (e Error) Code() Code {
	return e.code
}

func (e Error) Resp() (int, interface{}) {
	return e.code.status(), map[string]interface{}{
		"code":    e.code,
		"message": e.code.message(),
	}
}
