package main

func (e *Eval) startDefinition() {
	e.compiling = true
}

func (e *Eval) endDefinition() {
	e.Dict[e.tmp.Name] = e.tmp
	e.tmp = Word{Name: ""}
	e.compiling = false
}

func (e *Eval) leftParen() {
	e.comment++
}

func (e *Eval) store() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		xx := x.UnwrapVal()
		if xx < 0 {
			panic("!: STORE: Illegal attempt to change input")
		}
		e.Stack.items[xx] = y.UnwrapVal()
	}
}

func (e *Eval) rShift() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() >> x.UnwrapVal())
	}
}
