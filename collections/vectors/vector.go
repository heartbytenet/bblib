package vectors

var (
	VectorSizeDefault = 1024
)

type Vector struct {
	size int
	curr int
	body []byte
}

func (vector *Vector) Init() *Vector {
	vector.body = make([]byte, VectorSizeDefault)
	vector.size = VectorSizeDefault
	vector.curr = 0

	return vector
}

func (vector *Vector) Len() int {
	return vector.curr
}

func (vector *Vector) Size() int {
	return vector.size
}

func (vector *Vector) Extend() (size int) {
	var (
		data []byte
	)

	size = vector.size * 2
	data = make([]byte, size)

	vector.curr = copy(data, vector.body)
	vector.body = data
	vector.size = size

	return
}

func (vector *Vector) Write(data []byte) (size int, err error) {
	size = len(data)

	if (vector.curr + size) >= vector.size {
		vector.Extend()
	}

	size = copy(vector.body[vector.curr:], data)
	vector.curr += size

	return
}

func (vector *Vector) ReadAt(index int) byte {
	if index >= vector.size {
		return 0
	}

	return vector.body[index]
}

func (vector *Vector) ReadAll() (data []byte) {
	data = vector.body[:vector.curr]

	return
}

func (vector *Vector) Consume(size int) (data []byte) {

	if size >= vector.curr {
		data = make([]byte, vector.curr)
		copy(data, vector.body)

		vector.body = make([]byte, VectorSizeDefault)
		vector.size = VectorSizeDefault
		vector.curr = 0

		return
	}

	vector.curr -= size

	if vector.curr <= vector.size / 2 && vector.size / 2 >= VectorSizeDefault {
		vector.size /= 2
	}

	new_vector := make([]byte, vector.curr)
	copy(new_vector, vector.body[size:])

	data = make([]byte, size)
	copy(data, vector.body[:size])

	vector.body = make([]byte, vector.size)
	copy(vector.body, new_vector)

	return
}

func (vector *Vector) ConsumeWhen(fn func(*Vector) int) (size int, data []byte) {
	size = fn(vector)
	if size < 1 {
		return
	}

	data = vector.Consume(size)
	return
}
