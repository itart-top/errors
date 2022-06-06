package coder

import (
	"errors"
	"fmt"
	"log"
)

// 异常码
type withCode struct {
	C   Code
	Err error
}

func (ec *withCode) Error() string {
	return fmt.Sprintf("[%d] %s", ec.Code(), ec.Err.Error())
}

func (ec *withCode) Code() Code {
	return ec.C
}

func (ec *withCode) Format(state fmt.State, verb rune) {
	fmt.Println(string(verb))
	_, err := state.Write([]byte(ec.Error()))
	if err != nil {
		log.Fatal("write err: ", err)
	}
	var formatter fmt.Formatter
	if errors.As(ec.Err, &formatter) {
		formatter.Format(state, verb)
	}
}

func (ec *withCode) Is(err error) bool {
	if err == ec.C { // nolint
		return true
	}
	if x, ok := ec.Err.(interface{ Is(error) bool }); ok && x.Is(ec.Err) { // nolint:errorlint
		return true
	}

	return false
}

// 获取异常
func (ec *withCode) Unwrap() error {
	return ec.Err
}
