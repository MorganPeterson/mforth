package main

import "fmt"

func (e *Eval) plus() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() + x.UnwrapVal())
	}
}

func (e *Eval) minus() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() - x.UnwrapVal())
	}
}

func (e *Eval) star() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() * x.UnwrapVal())
	}
}

func (e *Eval) slash() {
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

func (e *Eval) starSlashMod() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	z := e.Stack.Pop()

	if x.IsOk() && y.IsOk() && z.IsOk() {
		xx := x.UnwrapVal()
		if xx == 0 {
			return
		}
		a := z.UnwrapVal() * y.UnwrapVal()
		e.Stack.Push(a % xx)
		e.Stack.Push(a / xx)
	}
}

func (e *Eval) plusStore() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	
	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() + x.UnwrapVal())
	}
}
		
