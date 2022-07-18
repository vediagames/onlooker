package errutil

import (
	"fmt"
	"strings"
)

type Error struct {
	errs []error
}

func (e *Error) Add(err error) {
	e.errs = append(e.errs, err)
}

func (e Error) Err() error {
	if len(e.errs) == 0 {
		return nil
	}

	return e
}

func (e Error) Error() string {
	return e.String()
}

func (e Error) String() string {
	if len(e.errs) == 0 {
		return ""
	}

	errs := make([]string, 0, len(e.errs))

	for _, err := range e.errs {
		errs = append(errs, err.Error())
	}

	return fmt.Sprintf("%s", strings.Join(errs, ", "))
}
