package vectors

import (
	"fmt"
	"strings"
	"testing"
)

func checkSlicing(t *testing.T, vector *Vector, size int) {

	var (
		slice        []byte
		trueSize     int
		trueCapacity int
	)

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

	if vector.Len() <= trueCapacity/2 && trueCapacity/2 >= VectorSizeDefault {
		trueCapacity /= 2
	}

	if vector.Size() != trueCapacity {
		t.Errorf("Expected capacity : '%d', but got '%d'", trueCapacity, vector.Size())
	}

}

func TestConsume(t *testing.T) {
	vector := &Vector{}
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

func TestVector_Init(t *testing.T) {
	vector := &Vector{}
	vector.Init()
}

func TestVector_Write(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	data := []byte("bla bla bla bla")

	size, err := vector.Write(data)
	if err != nil {
		t.Fatal("failed at writing vector:", err)
	}

	if size != len(data) {
		t.Fatal("returned length differs")
	}
}

func TestVector_Write_Empty(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	data := []byte("")

	size, err := vector.Write(data)
	if err != nil {
		t.Fatal("failed at writing vector:", err)
	}

	if size != len(data) {
		t.Fatal("returned length differs")
	}
}

func TestVector_Write_Nil(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	size, err := vector.Write(nil)
	if err != nil {
		t.Fatal("failed at writing vector:", err)
	}

	if size != 0 {
		t.Fatal("this should be zero")
	}
}

func TestVector_ReadAt(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	if v := vector.ReadAt(10); v != 0 {
		t.Fatal("this should return zero, got", v, "instead")
	}

	s := []byte("hello world")
	_, _ = vector.Write(s)

	for i := range s {
		if vector.ReadAt(i) != s[i] {
			t.Fatalf("expected [%v], instead got [%v]", s[i], vector.ReadAt(i))
		}
	}
}

func TestVector_ReadAll(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	s := []byte("hello world")
	_, _ = vector.Write(s)
	r := vector.ReadAll()

	for i := range s {
		if r[i] != s[i] {
			t.Fatalf("expected [%v], instead got [%v]", s[i], r[i])
		}
	}
}

func TestVector_Consume(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	s := []byte(strings.Repeat("hello world", 1024*16))
	_, _ = vector.Write(s)

	for i := range s {
		x := vector.Consume(1)[0]
		y := s[i]

		if x != y {
			t.Fatal("this should not fail")
		}
	}

	r := vector.Consume(len(s))
	if len(r) != 0 {
		t.Fatal("slice should be empty, instead got", r)
	}
}

func TestVector_Consume_Range(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	s := []byte(strings.Repeat("12345678", 1024))
	_, _ = vector.Write(s)

	for i := 0; i < 1024; i++ {
		x := vector.Consume(8)
		y := s

		for j := range x {
			if x[j] != y[j] {
				t.Fatal("this should not fail")
			}
		}
	}

	r := vector.Consume(len(s))
	if len(r) != 0 {
		t.Fatal("slice should be empty, instead got", r)
	}
}

func TestVector_ReadFrom(t *testing.T) {
	vector := &Vector{}
	vector.Init()

	data := []byte("wxyzabcd")
	_, _ = vector.Write(data)

	if vector.ReadFrom(-1, 2) != nil {
		t.Fatal("this should be nil")
	}

	if vector.ReadFrom(0, 0) != nil {
		t.Fatal("this should be nil")
	}

	if vector.ReadFrom(0, 10) != nil {
		t.Fatal("this should be nil")
	}

	for i, v := range vector.ReadFrom(0, 4) {
		if data[i] != v {
			t.Fatal("this should be equal:", v, data[i])
		}
	}
}
