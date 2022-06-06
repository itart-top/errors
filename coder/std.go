package coder

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const Unknown Code = -1

var out io.Writer = os.Stdout

func Log(target error) {
	_, err := out.Write([]byte(target.Error()))
	if err != nil {
		log.Fatal(err)
	}
	var stackErr *withStack
	if errors.As(target, &stackErr) {
		_, err = out.Write([]byte(fmt.Sprintf("%+v\n", stackErr.stack)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

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

func SetOutput(o io.Writer) {
	out = o
}
