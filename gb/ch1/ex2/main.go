package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("arg no %d is %s\n", idx, arg)
	}
}
