/**
 * @Title
 * @Description
 * @Author hyman
 * @Date 2022-06-05
 **/
package test

import (
	"fmt"
	"io"
	"testing"

	"github.com/itart-top/errors"
	"github.com/itart-top/errors/coder"
)

func TestWithMessagef(t *testing.T) {
	err := bizErr()
	//fmt.Println("----------")
	//fmt.Println(err)
	//fmt.Println("----------")
	//fmt.Println(fmt.Sprintf("%+v", err))
	//fmt.Println("----------")
	//fmt.Println(err.Error())
	//log.WithField("data", map[string]string{"ak":"av", "bk":"bv"}).Error(err)
	fmt.Println(fmt.Sprintf("%v", err))
}

func eof() error {
	return errors.WithMessage(io.EOF, "eor")
}

func bizErr() error {
	return coder.Code(110).WithMessage("bizerror")
}

func callControl() error {
	err := callService()
	return errors.WithMessage(err, "call service error")
}

func callService() error {
	err := callDao()
	return errors.WithMessage(err, "call dao error")
}

func callDao() error {
	return coder.Unknown.WithMessage("未知异常")
}
