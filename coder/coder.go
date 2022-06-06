package coder

import (
	E "errors"
	"fmt"

	"github.com/itart-top/errors"
)

type Code int

func (c Code) WithErrMessage(err error, msg string) error {
	if err == nil {
		return nil
	}
	err = &withCode{
		C:   c,
		Err: errors.WithMessage(err, msg),
	}
	if c > 0 {
		return err
	}
	return &withStack{
		error: err,
		stack: callers(),
	}
}
func (c Code) WithErrMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	err = &withCode{
		C:   c,
		Err: errors.WithMessagef(err, format, args),
	}
	if c > 0 {
		return err
	}
	return &withStack{
		error: err,
		stack: callers(),
	}
}

func (c Code) WithMessage(message string) error {
	err := &withCode{
		C:   c,
		Err: E.New(message),
	}
	if c > 0 {
		return err
	}
	return &withStack{
		error: err,
		stack: callers(),
	}
}

func (c Code) WithMessagef(format string, args ...interface{}) error {
	err := &withCode{
		C:   c,
		Err: fmt.Errorf(format, args...),
	}
	if c > 0 {
		return err
	}
	return &withStack{
		error: err,
		stack: callers(),
	}
}

func (c Code) Error() string {
	return fmt.Sprintf("[%d]", c)
}

func (c Code) WithErr(err error) error {
	if err == nil {
		return nil
	}
	err = &withCode{
		C:   c,
		Err: err,
	}
	if c > 0 {
		return err
	}
	return &withStack{
		error: err,
		stack: callers(),
	}
}
