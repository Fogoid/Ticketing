package queue

type Queue[T any] struct {
	count    int
	elements []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		count:    0,
		elements: make([]T, 0),
	}
}

func (q *Queue[T]) Push(element T) {
	q.elements = append(q.elements, element)
	q.count++
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.count <= 0 {
        var zeroVal T
		return zeroVal, false
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	q.count--

	return element, true
}

func (q *Queue[T]) IsEmpty() bool {
    return q.count == 0
}
