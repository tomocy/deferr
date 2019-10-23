package deferr_test

import (
	"errors"
	"fmt"

	"github.com/tomocy/deferr"
)

func ExampleFormat() {
	err := func() (err error) {
		defer deferr.Format(&err, "failed to foo")

		return errors.New("bar")
	}()

	fmt.Println(err)
	// Output:
	// failed to foo: bar
}
