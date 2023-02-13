package impl

type List[T any] interface {
	Add(i int, value T)
}

type ArrayList[T any] struct {
	Data []T
	Size int
}

func (a *ArrayList[T]) Add(i int, value T) {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Resize() {
	newData := make([]T, 0, a.Size*2)
	for i := range a.Data {
		newData[i] = a.Data[i]
	}
	a.Data = newData
}

var _ List[int] = &ArrayList[int]{}
