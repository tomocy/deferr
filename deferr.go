package deferr

import (
	"fmt"
)

func Format(err *error, format string, as ...interface{}) {
	if *err != nil {
		format, as = fmt.Sprintf("%s: %%v", format), append(as, *err)
		*err = fmt.Errorf(format, as...)
	}
}

func Wrapf(err *error, format string, as ...interface{}) {
	if *err != nil {
		format, as = fmt.Sprintf("%s: %%w", format), append(as, *err)
		*err = fmt.Errorf(format, as...)
	}
}

type Verb struct {
	flag, width, prec int
	verb              rune
}
