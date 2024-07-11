package sync

import (
	"github.com/heartbytenet/bblib/debug"
	"log"
	"reflect"
	"sync"
)

type Locked[T any] struct {
	value T
	mutex sync.Mutex
}

func NewLocked[T any](value T) *Locked[T] {
	return &Locked[T]{
		value: value,
		mutex: sync.Mutex{},
	}
}

func (locked *Locked[T]) Lock() {
	if debug.DEBUG {
		log.Println("locked lock", "t:", reflect.TypeOf(locked.value), "v:", locked.value)
	}

	locked.mutex.Lock()
}

func (locked *Locked[T]) Unlock() {
	if debug.DEBUG {
		log.Println("mutex unlock", "t:", reflect.TypeOf(locked.value), "v:", locked.value)
	}

	locked.mutex.Unlock()
}

func (locked *Locked[T]) Set(value T) {
	locked.Lock()
	defer locked.Unlock()

	locked.value = value
}

func (locked *Locked[T]) Get() (value T) {
	locked.Lock()
	defer locked.Unlock()

	value = locked.value

	return
}

func (locked *Locked[T]) GetForce() (value T) {
	value = locked.value

	return
}

func (locked *Locked[T]) Map(fn func(T) T) {
	if fn == nil {
		return
	}

	locked.Lock()
	defer locked.Unlock()

	locked.value = fn(locked.value)
}

func (locked *Locked[T]) Apply(fn func(T)) {
	if fn == nil {
		return
	}

	locked.Lock()
	defer locked.Unlock()

	fn(locked.value)
}
