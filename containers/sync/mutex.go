package sync

import "sync"

type Mutex[T any] struct {
	value T
	mutex sync.Mutex
}

func NewMutex[T any](value T) *Mutex[T] {
	return &Mutex[T]{
		value: value,
		mutex: sync.Mutex{},
	}
}

func (mutex *Mutex[T]) Lock() {
	mutex.mutex.Lock()
}

func (mutex *Mutex[T]) Unlock() {
	mutex.mutex.Unlock()
}

func (mutex *Mutex[T]) Set(value T) {
	mutex.Lock()
	defer mutex.Unlock()

	mutex.value = value
}

func (mutex *Mutex[T]) Get() (value T) {
	mutex.Lock()
	defer mutex.Unlock()

	value = mutex.value

	return
}

func (mutex *Mutex[T]) Map(fn func(T) T) {
	if fn == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	mutex.value = fn(mutex.value)
}

func (mutex *Mutex[T]) Apply(fn func(T)) {
	if fn == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	fn(mutex.value)
}
