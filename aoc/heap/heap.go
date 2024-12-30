package heap

import sheap "container/heap"

type Heap[T any] struct {
	impl heapImpl[T]
}

func New[T any](less func(T, T) bool) *Heap[T] {
	return &Heap[T]{impl: heapImpl[T]{less: less}}
}

func (pq *Heap[T]) Push(v T) {
	sheap.Push(&pq.impl, v)
}

func (pq *Heap[T]) Pop() T {
	v := sheap.Pop(&pq.impl)
	return v.(T)
}

func (pq *Heap[T]) Len() int {
	return pq.impl.Len()
}

func (pq *Heap[T]) Empty() bool {
	return pq.Len() == 0
}

func (pq *Heap[T]) Peek() T {
	return pq.impl.items[0]
}

type heapImpl[T any] struct {
	items []T
	less  func(T, T) bool
}

func (w heapImpl[T]) Len() int { return len(w.items) }

func (w heapImpl[T]) Less(i, j int) bool { return w.less(w.items[i], w.items[j]) }

func (w heapImpl[T]) Swap(i, j int) { w.items[i], w.items[j] = w.items[j], w.items[i] }

func (w *heapImpl[T]) Pop() any {
	n := len(w.items)
	value := w.items[n-1]
	w.items = w.items[:n-1]

	return value
}

func (w *heapImpl[T]) Push(v any) {
	w.items = append(w.items, v.(T))
}
