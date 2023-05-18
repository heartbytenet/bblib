package sync

import (
	"github.com/heartbytenet/bblib/debug"
	"log"
	"reflect"
	"sync"
)

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
	if debug.DEBUG {
		log.Println("mutex lock", "t:", reflect.TypeOf(mutex.value), "v:", mutex.value)
	}

	mutex.mutex.Lock()
}

func (mutex *Mutex[T]) Unlock() {
	if debug.DEBUG {
		log.Println("mutex unlock", "t:", reflect.TypeOf(mutex.value), "v:", mutex.value)
	}

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

func (mutex *Mutex[T]) GetForce() (value T) {
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
