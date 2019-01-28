package apperr

import "fmt"

var (
	NotFound             = "NotFound"
	BindError            = "BindError"
	DBError              = "DBError"
	ExternalServiceError = "ExternalServiceError"
	ServerError          = "ServerError"
	BadRequest           = "BadRequest"
	AuthorizationError   = "AuthorizationError"
	ValidationError      = "ValidationError"
)

// Error is アプリケーション内で扱われる標準的なエラー
type Error struct {
	code    string
	origErr error
}

func NewError(code string, err error) *Error {
	return &Error{
		code:    code,
		origErr: err,
	}
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.origErr.Error())
}

func (e *Error) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"code":  e.code,
		"error": e.origErr.Error(),
	}
}
