package main

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
	if x != nil && y != nil {
		e.Stack.Push(*y / *x)
	}
}

