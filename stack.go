package main

import (
	"fmt"
)

var (
	ErrStackUnderflow = "stack underflow"
)

type Stack[T any] struct {
	items []T
	length int64 `default:0`
}

func (s *Stack[T]) IsEmpty() bool {
	return (*s).length == 0
}

func (s *Stack[T]) Push(x T) {
	(*s).items = append((*s).items, x)
	(*s).length++
}

func (s *Stack[T]) Pop() Result[T] {
	if s.IsEmpty() {
		return Error[T](fmt.Errorf(ErrStackUnderflow))
	}

	i := (*s).length - 1
	x := (*s).items[i]
	(*s).items = (*s).items[:i]
	(*s).length--

	return Ok[T](x)
}

func (s *Stack[T]) Peek() Result[T] {
	if s.IsEmpty() {
		return Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	i := (*s).length - 1
	if i < 0 {
		return Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	return Ok[T]((*s).items[i])
}

func (s *Stack[T]) Fetch(item int64) Result[T] {
	if s.IsEmpty() {
		return Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	i := (*s).length - item - 1
	if i < 0 {
		return Error[T](fmt.Errorf(ErrStackUnderflow))
	}
	return Ok[T]((*s).items[i])
}	

func (e *Eval) dot() {
	x := e.Stack.Pop()
	if x.IsOk() {
		fmt.Printf("%d\n", x.UnwrapVal())
		return
	}
	fmt.Printf("%s\n", x.UnwrapErr())
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

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(x.UnwrapVal())
		e.Stack.Push(y.UnwrapVal())
	}
}

func (e *Eval) twoSwap() {
	v := e.Stack.Pop()
	w := e.Stack.Pop()
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if v.IsOk() && w.IsOk() && x.IsOk() && y.IsOk() {
		e.Stack.Push(w.UnwrapVal())
		e.Stack.Push(v.UnwrapVal())
		e.Stack.Push(y.UnwrapVal())
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) dup() {
	x := e.Stack.Peek()
	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) twoDup() {
	x := e.Stack.Peek()
	y := e.Stack.Fetch(2)
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal())
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) nonZeroDup() {
	x := e.Stack.Peek()
	if x.IsOk() {
		xx := x.UnwrapVal()
		if xx != 0 {
			e.Stack.Push(xx)
		}
	}
}

func (e *Eval) over() {
	x := e.Stack.Fetch(e.Stack.length - 2)
	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) twoOver() {
	x := e.Stack.Fetch(3)
	y := e.Stack.Fetch(4)

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal())
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) pick() {
	x := e.Stack.Pop()
	if x.IsOk() {
		xx := x.UnwrapVal() + 1
		y := e.Stack.Fetch(e.Stack.length - xx)
		if y.IsOk() {
			e.Stack.Push(y.UnwrapVal())
		}
	}
}

func (e *Eval) rot() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	z := e.Stack.Pop()
	if x.IsOk() && y.IsOk() && z.IsOk() {
		e.Stack.Push(y.UnwrapVal())
		e.Stack.Push(x.UnwrapVal())
		e.Stack.Push(z.UnwrapVal())
	}
}

func (e *Eval) reverseRot() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	z := e.Stack.Pop()
	if x.IsOk() && y.IsOk() && z.IsOk() {
		e.Stack.Push(x.UnwrapVal())
		e.Stack.Push(z.UnwrapVal())
		e.Stack.Push(y.UnwrapVal())
	}
}

func (e *Eval) depth() {
	e.Stack.Push(e.Stack.length)
}

func (e *Eval) roll() {
	x := e.Stack.Pop()

	if x.IsOk() {
		xx := x.UnwrapVal()
		switch xx {
		case 0:
			return
		case 1:
			e.swap()
		case 2:
			e.rot()
		default:
			i := e.Stack.length - (xx + 1)
			y := e.Stack.items[i]
			copy(e.Stack.items[i:], e.Stack.items[i+1:])
			e.Stack.items[e.Stack.length-1] = y
		}
	}
}

/*
 * return stack functions
 */

func (e *Eval) toR() {
	x := e.Stack.Pop()
	if x.IsOk() {
		e.RStack.Push(x.UnwrapVal())
	}
}

func (e *Eval) fromR() {
	x := e.RStack.Pop()
	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) fetchR() {
	x := e.RStack.Peek()
	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) twoToR() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x.IsOk() && y.IsOk() {
		e.RStack.Push(y.UnwrapVal())
		e.RStack.Push(x.UnwrapVal())
	}
}

func (e *Eval) twoFromR() {
	x := e.RStack.Pop()
	y := e.RStack.Pop()
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal())
		e.Stack.Push(x.UnwrapVal())
	}
}

func (e *Eval) fetchTwoR() {
	x := e.RStack.Fetch(2)
	y := e.RStack.Peek()
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(x.UnwrapVal())
		e.Stack.Push(y.UnwrapVal())
	}
}
