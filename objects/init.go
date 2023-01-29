package objects

type Initializable[T any] interface {
	Init() *T
}

func Init[T any](obj Initializable[T]) *T {
	return obj.Init()
}
