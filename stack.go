package main

import (
	"fmt"
)

type Stack[T interface{}] struct {
	items []T
	length int64
}

func (s *Stack[T]) IsEmpty() bool {
	return (*s).length == 0
}

func (s *Stack[T]) Push(x T) {
	(*s).items = append((*s).items, x)
	(*s).length++
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		fmt.Printf("stack underflow\n")
		return nil
	}

	i := (*s).length - 1
	x := (*s).items[i]
	(*s).items = (*s).items[:i]
	(*s).length--

	return &x
}

func (s *Stack[T]) Peek() *T {
	if s.IsEmpty() {
		fmt.Printf("stack underflow\n")
		return nil
	}

	i := (*s).length - 1
	if i < 0 {
		fmt.Printf("stack underflow\n")
		return nil
	}
	return &(*s).items[i]
}

func (s *Stack[T]) Fetch(item int64) *T {
	if s.IsEmpty() {
		fmt.Printf("stack underflow\n")
		return nil
	}
	i := (*s).length - 2
	if i < 0 {
		fmt.Printf("stack underflow\n")
		return nil
	}
	return &(*s).items[(*s).length - 2]
}	

func (e *Eval) print() {
	x := e.Stack.Pop()
	if x != nil {
		fmt.Printf("%d\n", *x)
	}
}

func (e *Eval) drop() {
	e.Stack.Pop()
}

func (e *Eval) twoDrop() {
	e.Stack.Pop()
	e.Stack.Pop()
}

func (e *Eval) swap() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x != nil && y != nil {
		e.Stack.Push(*x)
		e.Stack.Push(*y)
	}
}

func (e *Eval) twoSwap() {
	v := e.Stack.Pop()
	w := e.Stack.Pop()
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if v != nil && w != nil && x != nil && y != nil {
		e.Stack.Push(*w)
		e.Stack.Push(*v)
		e.Stack.Push(*y)
		e.Stack.Push(*x)
	}
}

func (e *Eval) dup() {
	x := e.Stack.Peek()
	if x != nil {
		e.Stack.Push(*x)
	}
}

func (e *Eval) twoDup() {
	x := e.Stack.Peek()
	y := e.Stack.Fetch(2)
	if x != nil && y != nil {
		e.Stack.Push(*y)
		e.Stack.Push(*x)
	}
}

func (e *Eval) nonZeroDup() {
	x := e.Stack.Peek()
	if x != nil && *x != 0 {
		e.Stack.Push(*x)
	}
}

func (e *Eval) over() {
	x := e.Stack.Fetch(e.Stack.length - 2)
	e.Stack.Push(*x)
}

func (e *Eval) twoOver() {
	x := e.Stack.Fetch(3)
	y := e.Stack.Fetch(4)
	e.Stack.Push(*y)
	e.Stack.Push(*x)
}

func (e *Eval) pick() {
	x := e.Stack.Pop()
	if x != nil {
		*x++
		y := e.Stack.Fetch(e.Stack.length - 2)
		if y != nil {
			e.Stack.Push(*y)
		}
	}
}

func (e *Eval) rot() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	z := e.Stack.Pop()
	if x != nil && y != nil && z != nil {
		e.Stack.Push(*y)
		e.Stack.Push(*x)
		e.Stack.Push(*z)
	}
}

func (e *Eval) reverseRot() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	z := e.Stack.Pop()
	if x != nil && y != nil && z != nil {
		e.Stack.Push(*x)
		e.Stack.Push(*z)
		e.Stack.Push(*y)
	}
}

func (e *Eval) depth() {
	e.Stack.Push(e.Stack.length)
}

func (e *Eval) roll() {
	x := e.Stack.Pop()
	if x == nil {
		return
	}

	switch *x {
	case 0:
		return
	case 1:
		e.swap()
	case 2:
		e.rot()
	default:
		i := (*e).Stack.length - (*x + 1)
		y := e.Stack.items[i]
		copy(e.Stack.items[i:], e.Stack.items[i+1:])
		e.Stack.items[e.Stack.length-1] = y
	}
}

/*
 * return stack functions
 */

func (e *Eval) toR() {
	x := e.Stack.Pop()
	if x != nil {
		e.RStack.Push(*x)
	}
}

func (e *Eval) fromR() {
	x := e.RStack.Pop()
	if x != nil {
		e.Stack.Push(*x)
	}
}

func (e *Eval) fetchR() {
	x := e.RStack.Peek()
	if x != nil {
		e.Stack.Push(*x)
	}
}

func (e *Eval) twoToR() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x != nil && y != nil {
		e.RStack.Push(*y)
		e.RStack.Push(*x)
	}
}

func (e *Eval) twoFromR() {
	x := e.RStack.Pop()
	y := e.RStack.Pop()
	if x != nil && y != nil {
		e.Stack.Push(*y)
		e.Stack.Push(*x)
	}
}

func (e *Eval) fetchTwoR() {
	x := e.RStack.Fetch(2)
	y := e.RStack.Peek()
	if x != nil && y != nil {
		e.Stack.Push(*x)
		e.Stack.Push(*y)
	}
}
