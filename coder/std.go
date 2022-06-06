package coder

import (
	"errors"
)

const Unknown Code = -1

func Case(target error) error {
	err := target
	i := 0
	for {
		raw := errors.Unwrap(err)
		if raw == nil {
			return err
		}
		err = raw
		i++
		if i == 20 { // 防止死循环
			return target
		}
	}
}

func Extract(err error) Code {
	var code *withCode
	if errors.As(err, &code) {
		return code.Code()
	}
	return Unknown
}
