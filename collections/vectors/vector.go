package vectors

var (
	VectorSizeDefault = 64
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

	copy(data, vector.body)

	vector.body = data
	vector.size = size

	return
}

func (vector *Vector) Append(v byte) {
	for (vector.curr + 1) >= vector.size {
		vector.Extend()
	}

	vector.body[vector.curr] = v
	vector.curr++

	return
}

func (vector *Vector) Write(data []byte) (size int, err error) {
	for _, v := range data {
		vector.Append(v)
	}

	size = len(data)
	return
}

func (vector *Vector) ReadAt(index int) byte {
	if index >= vector.size {
		return 0
	}

	return vector.body[index]
}

func (vector *Vector) ReadFrom(index int, length int) (data []byte) {
	if index < 0 {
		return
	}

	if index == length {
		return
	}

	if (index + length) >= vector.Len() {
		return
	}

	data = vector.body[index : index+length][:]
	return
}

func (vector *Vector) ReadAll() (data []byte) {
	data = vector.body[:vector.curr][:]

	return
}

func (vector *Vector) Consume(size int) (data []byte) {

	var (
		body []byte
	)

	if size >= vector.curr {
		data = vector.body[:vector.curr]

		vector.body = make([]byte, VectorSizeDefault)
		vector.size = VectorSizeDefault
		vector.curr = 0

		return
	}

	vector.curr -= size

	if vector.curr <= vector.size/2 && vector.size/2 >= VectorSizeDefault {
		vector.size /= 2
	}

	data = vector.body[:size]
	body = vector.body[size:]

	vector.body = make([]byte, vector.size)
	copy(vector.body, body)

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
