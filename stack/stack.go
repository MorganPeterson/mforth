package stack

import (
	"fmt"

	"github.com/MorganPeterson/mForth/result"
)

var (
	ErrStackUnderflow = "stack underflow"
)

type Stack[T any] struct {
	items []T
	length int `default:0`
}

func (s *Stack[T]) IsEmpty() bool {
	return (*s).length == 0
}

func (s *Stack[T]) Push(x T) {
	(*s).items = append((*s).items, x)
	(*s).length++
}

func (s *Stack[T]) Pop() result.Result[T] {
	if s.IsEmpty() {
		return result.Error[T](fmt.Errorf(ErrStackUnderflow))
	}

	i := (*s).length - 1
	x := (*s).items[i]
	(*s).items = (*s).items[:i]
	(*s).length--

	return result.Ok[T](x)
}

func (s *Stack[T]) Peek() result.Result[T] {
	if s.IsEmpty() {
		return result.Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	i := (*s).length - 1
	if i < 0 {
		return result.Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	return result.Ok[T]((*s).items[i])
}

func (s *Stack[T]) Fetch(item int) result.Result[T] {
	if s.IsEmpty() {
		return result.Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	i := (*s).length - item - 1
	if i < 0 {
		return result.Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	return result.Ok[T]((*s).items[i])
}	

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) Insert(index int, val T) {
	(*s).items[index] = val
}

func (s *Stack[T]) Rm(index int) {
	copy(s.items[index:], s.items[index+1:])
}

func (s *Stack[T]) PeekAt(index int) T {
	return (*s).items[index]
}
