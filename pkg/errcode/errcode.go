package errcode

import (
	"fmt"
)

type Error struct {
	code int
	msg  string
	data []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg

	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Data() []string {
	return e.data
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.data = []string{}

	for _, data := range details {
		newError.data = append(newError.data, data)
	}

	return &newError
}
