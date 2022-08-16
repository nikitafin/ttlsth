package main

import "fmt"

func main() {
	sl := make([]int, 100)
	for i := range sl {
		sl[i] = i + 1
	}
	sl = append(sl[:10], sl[len(sl)-10:]...)
	for i := 0; i < len(sl) / 2; i++ {
		sl[i], sl[len(sl)-i-1] = sl[len(sl)-i-1], sl[i]
	}
	fmt.Println(sl)
}
