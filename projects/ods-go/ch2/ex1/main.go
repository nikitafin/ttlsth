package main

import (
	"impl"
)

type ArrayListWithAddAll[T any] struct {
	impl.ArrayList[T]
}

func (r *ArrayListWithAddAll[T]) resize(count int) {
	newData := make([]T, 0, r.Size*2+count)
	for i := range r.Data {
		newData[i] = r.Data[i]
	}
	r.Data = newData
}

func (r *ArrayListWithAddAll[T]) AddAll(pos uint, c ...T) {
	r.resize(len(c))
}

func main() {
	a := ArrayListWithAddAll[int]{}
	a.Add(0, 10)
}
