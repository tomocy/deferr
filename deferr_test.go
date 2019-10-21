package deferr

import (
	"errors"
	"testing"
)

func TestFormat(t *testing.T) {
	expected := "failed to foo: bar"
	actual := func() (err error) {
		defer Format(&err, "failed to foo")

		return errors.New("bar")
	}()

	if actual.Error() != expected {
		t.Errorf("unexpected error: got %s, expected %s\n", actual, expected)
	}
}
