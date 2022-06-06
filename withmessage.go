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
	"log"
)

type withMessage struct {
	err error
	msg string
}

func (w *withMessage) Error() string { return w.msg + " -> " + w.err.Error() }
func (w *withMessage) Cause() error  { return w.err }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withMessage) Unwrap() error { return w.err }

func (w *withMessage) Format(state fmt.State, verb rune) {

	_, err := state.Write([]byte(w.msg + " -> "))
	if err != nil {
		log.Fatal("write err: ", err)
	}

	var formatter fmt.Formatter

	if errors.As(w.err, &formatter) {
		formatter.Format(state, verb)
	} else {
		_, err := state.Write([]byte(w.err.Error()))
		if err != nil {
			log.Fatal("write err: ", err)
		}
	}

}
