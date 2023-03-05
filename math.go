package main

import "fmt"

func (e *Eval) add() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() + x.UnwrapVal())
	}
}

func (e *Eval) sub() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() - x.UnwrapVal())
	}
}

func (e *Eval) mul() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() * x.UnwrapVal())
	}
}

func (e *Eval) div() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x.IsOk() && y.IsOk() {
		xx := x.UnwrapVal()
		if xx == 0 {
			fmt.Printf("/: zero division error\n")
			return
		}
		e.Stack.Push(y.UnwrapVal() / xx)
	}
}

func (e *Eval) mod() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x.IsOk() || y.IsOk() {
		xx := x.UnwrapVal()
		if xx == 0 {
			fmt.Printf("MOD: zero division error\n")
			return
		}
		e.Stack.Push(y.UnwrapVal() % xx)
	}
}

func (e *Eval) twoStar() {
	x := e.Stack.Pop()

	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal() << 1)
	}
}

func (e *Eval) twoSlash() {
	x := e.Stack.Pop()

	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal() >> 1)
	}
}
