package main

import (
	"basic_go_practicum/pkg/stopwatch"
	"fmt"
	"time"
)

func main() {
	sw := stopwatch.Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
