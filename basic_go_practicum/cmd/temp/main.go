package main

import "fmt"

func main() {
	i := 10
	defer func(i *int) {
		*i = *i + 10
		fmt.Println("Defer", *i)
	}(&i)
	fmt.Println(i)
	i = i + 100
	fmt.Println(i)
}
