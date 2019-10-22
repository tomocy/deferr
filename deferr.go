package deferr

import (
	"fmt"
	"sort"
	"strings"
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

func (v Verb) String() string {
	var w strings.Builder
	fmt.Fprint(&w, "%")
	if isFlag(v.flag) {
		fmt.Fprint(&w, string([]rune{rune(v.flag)}))
	}
	if v.width != 0 {
		fmt.Fprint(&w, v.width)
	}
	if v.prec != 0 {
		fmt.Fprint(&w, ".", v.prec)
	}
	fmt.Fprint(&w, string([]rune{v.verb}))

	return w.String()
}

func isFlag(f int) bool {
	sort.Ints(flags)
	i := sort.Search(len(flags), func(i int) bool {
		return f <= flags[i]
	})

	return i < len(flags) && flags[i] == f
}

var flags = []int{' ', '#', '+', '-', '0'}
