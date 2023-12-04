package utils

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	Len  int
}

func (q *Queue[T]) Enqueue(value T) {
	n := node[T]{value: value, next: nil}

	if q.Len == 0 {
		q.head = &n
		q.tail = &n
		q.Len++
		return
	}

	q.tail.next = &n
	q.tail = &n
	q.Len++
}

func (q *Queue[T]) Dequeue() (value *T, ok bool) {
	if q.Len == 0 {
		return nil, false
	}

	head := q.head
	q.head = q.head.next
	head.next = nil
	q.Len--

	return &head.value, true
}
