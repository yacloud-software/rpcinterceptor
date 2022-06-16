package evaluators

import (
	"fmt"
)

type logger struct {
	prefix string
}

func newLogger() *logger {
	res := &logger{}
	res.prefix = "[evaluator] "
	return res
}

func (l *logger) Printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	fmt.Printf("%s%s", l.prefix, s)
}
