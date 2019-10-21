package deferr

import (
	"errors"
	"fmt"
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
	if errors.Unwrap(actual) != nil {
		t.Errorf("unexpected unwrapeed error: got %s, expected nil", errors.Unwrap(actual))
	}
}

func TestWrapf(t *testing.T) {
	base := errors.New("bar")
	expected := fmt.Errorf("failed to foo: %w", base)
	actual := func() (err error) {
		defer Wrapf(&err, "failed to foo")

		return errors.New("bar")
	}()

	if actual.Error() != expected.Error() {
		t.Errorf("unexpected error: got %s, expected %s", actual, expected)
	}
	if errors.Unwrap(actual).Error() != base.Error() {
		t.Errorf("unexpected unwrapped error: got %s, expected %s", errors.Unwrap(actual), base)
	}
}
