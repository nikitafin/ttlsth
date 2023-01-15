package my_check

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrFoundNumbers  = errors.New("found numbers")
	ErrLineIsTooLong = errors.New("line is too long")
	ErrNoTwoSpaces   = errors.New("no two spaces")
)

type Errs []error

func (e Errs) Error() string {
	return fmt.Sprintf("%v", e)
}

func MyCheck(input string) error {
	var errs Errs
	if strings.ContainsAny(input, "0123456789") {
		errs = append(errs, ErrFoundNumbers)
	}
	if len(input) > 20 {
		errs = append(errs, ErrLineIsTooLong)
	}
	if strings.Count(input, " ") < 2 {
		errs = append(errs, ErrNoTwoSpaces)
	}
	return errs
}
