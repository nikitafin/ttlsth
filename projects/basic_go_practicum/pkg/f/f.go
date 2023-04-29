package f

import "fmt"

func f(a, b, i int) {
	if i == 0 {
		panic("i == 0")
	}

	f(b, a/b, i-1)
}

func Test(a, b, i int) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()
	f(a, b, i)

	return err
}
