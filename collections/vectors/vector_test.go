package vectors

import (
	"fmt"
	"testing"
	"github.com/heartbytenet/bblib/collections/vectors"
)

func checkSlicing(t *testing.T, vector *vectors.Vector, size int) {

	var (
		slice []byte
		trueSize int
		trueCapacity int
		VectorSizeDefault int
	)

	VectorSizeDefault = vectors.VectorSizeDefault

	trueCapacity = vector.Size()
	trueSize = vector.Len() - size
	if trueSize < 0 {
		trueSize = 0
	}


	if size > len(vector.ReadAll()) {
		size = len(vector.ReadAll())
	}

	slice = vector.ReadAll()[:size]

	data := vector.Consume(size)

	if string(data) != string(slice) {
		t.Errorf("Expected sliced data : '%s', but got '%s'", string(slice), string(data))
		t.Errorf("In hexadecimal - Expected: %x, got: %x", slice, data)
	}

	if vector.Len() != trueSize {
		t.Errorf("Expected size : '%d', but got '%d'", trueSize, vector.Len())
	}

	if vector.Len() <= trueCapacity / 2 && trueCapacity / 2 >= VectorSizeDefault {
		trueCapacity /= 2
	}

	if vector.Size() != trueCapacity {
		t.Errorf("Expected capacity : '%d', but got '%d'", trueCapacity, vector.Size())
	}

}

func TestConsume(t *testing.T) {
	vector := &vectors.Vector{}
	vector.Init()

	fmt.Println("-- TESTING BYTES CONSUMPTION --")

	vector.Write([]byte("aaaaaaaaaabbbbbbbbbbccccccccccdddddddddd"))

	t.Run("CheckSlicing | 10", func(t *testing.T) {
		checkSlicing(t, vector, 10)
	})

	t.Run("CheckSlicing | 5", func(t *testing.T) {
		checkSlicing(t, vector, 5)
	})

	t.Run("CheckSlicing | 20", func(t *testing.T) {
		checkSlicing(t, vector, 20)
	})

	t.Run("CheckSlicing | 50", func(t *testing.T) {
		checkSlicing(t, vector, 50)
	})

	fmt.Println("-- TESTING LARGE BYTES CONSUMPTION --")

	vector.Write([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

	t.Run("CheckSlicing | 100", func(t *testing.T) {
		checkSlicing(t, vector, 100)
	})

	t.Run("CheckSlicing | 10000", func(t *testing.T) {
		checkSlicing(t, vector, 10000)
	})
}