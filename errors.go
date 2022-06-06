/**
 * @Title
 * @Description
 * @Author hyman
 * @Date 2022-06-05
 **/
package errors

import (
	"errors"
	"fmt"
)

func WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		err: err,
		msg: err.Error() + ":" + fmt.Sprintf(format, args...),
	}
}

func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		err: err,
		msg: message,
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}
