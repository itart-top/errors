/**
 * @Title
 * @Description
 * @Author hyman
 * @Date 2022-06-06
 **/
package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/itart-top/errors/coder"
)

func Test_coder(t *testing.T) {
	t.Parallel()
	t.Run("test unexpected", func(t *testing.T) {
		expected := "[-1] db error"
		err := coder.Unknown.WithMessage("db error")
		if err.Error() != expected {
			t.Fatalf("execped = %s actual = %s", expected, err.Error())
		}
		stackInfo := fmt.Sprint(err)
		if !strings.Contains(stackInfo, "coder_test.go") {
			t.Fatalf("not contain stack info: %s", stackInfo)
		}
	})
}
