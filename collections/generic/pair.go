package generic

type Pair[T any, U any] struct {
	a T
	b U
}

func NewPair[T any, U any](a T, b U) Pair[T, U] {
	return Pair[T, U]{
		a,
		b,
	}
}

func (pair Pair[T, U]) A() T {
	return pair.a
}

func (pair Pair[T, U]) B() U {
	return pair.b
}
