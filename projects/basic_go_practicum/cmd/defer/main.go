package main

import "fmt"

var Global = 5

func UseGlobal() {
	defer func() {
		Global = 5
	}()
	Global = 10
	fmt.Println(Global)
}

func main() {
	UseGlobal()
	fmt.Println(Global)
}
