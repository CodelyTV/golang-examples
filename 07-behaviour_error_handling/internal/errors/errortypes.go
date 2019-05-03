package errors

import "fmt"

type BadResponseErr struct {
	Msg  string
	File string
	Line int
}

func (e *BadResponseErr) Error() string {
	return fmt.Sprintf("%s: %d: %s", e.File, e.Line, e.Msg)
}
