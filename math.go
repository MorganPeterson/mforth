package main

import "fmt"

func (e *Eval) add() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x != nil && y != nil {
		e.Stack.Push(*y + *x)
	}
}

func (e *Eval) sub() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x != nil && y != nil {
		e.Stack.Push(*y - *x)
	}
}

func (e *Eval) mul() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x != nil && y != nil {
		e.Stack.Push(*y * *x)
	}
}

func (e *Eval) div() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x == nil || y == nil {
		return
	}
	if *x == 0 {
		fmt.Printf("/: zero division error\n")
		return
	}
	e.Stack.Push(*y / *x)
}

func (e *Eval) mod() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x == nil || y == nil {
		return
	}

	if *x == 0 {
		fmt.Printf("MOD: zero division error\n")
		return
	}
	e.Stack.Push(*y % *x)
}

func (e *Eval) twoStar() {
	x := e.Stack.Pop()

	if x == nil {
		return
	}
	e.Stack.Push(*x << 1)
}

func (e *Eval) twoSlash() {
	x := e.Stack.Pop()

	if x == nil {
		return
	}
	e.Stack.Push(*x >> 1)
}
