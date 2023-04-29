package main

import (
	"basic_go_practicum/pkg/f"
	"fmt"
)

type input struct {
	a       int
	b       int
	counter int
}

func main() {
	for _, pars := range []input{
		{10, 5, 3},
		{100, 7, 10},
		{1, 1, 1000},
	} {
		fmt.Printf("(%d, %d, %d) => %v\r\n", pars.a, pars.b, pars.counter,
			f.Test(pars.a, pars.b, pars.counter))
	}
}
