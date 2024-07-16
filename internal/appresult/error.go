package appresult

import "encoding/json"

var (
	ErrMissingParam   = NewAppError(nil, "missing param", "SE-00001")
	ErrNotFound       = NewAppError(nil, "not found", "SE-00003")
	ErrInternalServer = NewAppError(nil, "internal server error", "SE-00004")
	ErrNotAcceptable  = NewAppError(nil, "Not Acceptable", "SE-00007")
	ErrFileSize       = NewAppError(nil, "Too Large", "SE-00012")

	UserNotExist = NewAppError(nil, "User Not Exist", "")

	PasswordIncorrect = NewAppError(nil, "Password Incorrect", "")
)

type AppError struct {
	Status  bool   `json:"status"`
	Err     error  `json:"-"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, code string) *AppError {
	return &AppError{
		Status:  false,
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "Internal Server Error", "SE-000")
}
