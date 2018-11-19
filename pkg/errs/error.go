package errs

import "fmt"

func New(err error, code, desc string) error {
	return &Error{
		Code: code,
		Desc: desc,
		Err:  err.Error(),
	}
}

type Error struct {
	Code string `json:"StatusCode"`
	Desc string `json:"StatusDesc"`
	Err  string `json:"error"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%s, %s", e.Code, e.Desc, e.Error)
}
