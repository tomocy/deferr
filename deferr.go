package deferr

import (
	"fmt"
	"sort"
	"strings"
)

func Format(err *error, format string, as ...interface{}) {
	DefaultVerbMap.Format(err, format, as...)
}

func Wrapf(err *error, format string, as ...interface{}) {
	DefaultVerbMap.Wrapf(err, format, as...)
}

var DefaultVerbMap = VerbMap{
	KeyFormat: {Verb: 'v'}, KeyWrap: {Verb: 'w'},
}

type VerbMap map[verbKey]Verb

func (m VerbMap) Format(err *error, format string, as ...interface{}) {
	if *err == nil {
		return
	}

	format, as = wrapFormat(format, m[KeyFormat]), append(as, *err)
	*err = fmt.Errorf(format, as...)
}

func (m VerbMap) Wrapf(err *error, format string, as ...interface{}) {
	if *err == nil {
		return
	}

	format, as = wrapFormat(format, m[KeyWrap]), append(as, *err)
	*err = fmt.Errorf(format, as...)
}

const (
	_ verbKey = iota
	KeyFormat
	KeyWrap
)

type verbKey int

func wrapFormat(format string, verb Verb) string {
	return fmt.Sprintf("%s: %s", format, verb)
}

type Verb struct {
	Flag, Width, Prec int
	Verb              rune
}

func (v Verb) String() string {
	var w strings.Builder
	fmt.Fprint(&w, "%")
	if isFlag(v.Flag) {
		fmt.Fprint(&w, string([]rune{rune(v.Flag)}))
	}
	if v.Width != 0 {
		fmt.Fprint(&w, v.Width)
	}
	if v.Prec != 0 {
		fmt.Fprint(&w, ".", v.Prec)
	}
	fmt.Fprint(&w, string([]rune{v.Verb}))

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
