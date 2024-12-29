package stack

type Stack[T any] struct {
	items []T
}

type Option[T any] func(*Stack[T])

func WithCapacity[T any](capacity int) func(*Stack[T]) {
	return func(s *Stack[T]) {
		s.items = make([]T, 0, capacity)
	}
}

func New[T any](options ...Option[T]) Stack[T] {
	s := Stack[T]{}
	for _, option := range options {
		option(&s)
	}

	return s
}

func (s *Stack[T]) Push(values ...T) {
	s.items = append(s.items, values...)
}

func (s *Stack[T]) Pop() T {
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return value
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
