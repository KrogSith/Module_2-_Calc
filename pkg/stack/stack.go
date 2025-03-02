package stack

type stack[T comparable] struct {
	array []T
}

// NewStack() Создаёт новый стек
func NewStack[T comparable]() *stack[T] { return &stack[T]{} }

// Push() Добавляет элемент в конец стека
func (s *stack[T]) Push(n T) {
	s.array = append(s.array, n)
}

// Pop() Удаляет последний элемент стека
func (s *stack[T]) Pop() T {
	element := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return element
}

// Len() Возвращает длину стека
func (s *stack[T]) Len() int {
	return len(s.array)
}

// GetArray() Возвращает стек
func (s *stack[T]) GetArray() []T {
	return s.array
}
