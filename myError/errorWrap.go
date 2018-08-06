package MyError

import (
	//"errors"

	"github.com/pkg/errors"
)

func WrapTest(a, b int) string {
	err := errors.New("a")
	if b == 0 {
		err = errors.Wrap(err, "b")
	}
	return err.Error()
}
