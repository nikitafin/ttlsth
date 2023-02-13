package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// files := make([]string)
	var files []string
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't open file: %s. Err: %v", filename, err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 {
				files = append(files, filename)
			}
		}
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
