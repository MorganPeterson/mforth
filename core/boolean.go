package core

var (
	TT = -1
	FF = 0
)

func If(cond bool) int {
	if cond {
		return TT
	}
	return FF
}

func (e *Eval) ftrue() {
	e.Stack.Push(TT)
}

func (e *Eval) ffalse() {
	e.Stack.Push(FF)
}

func (e *Eval) fand() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() & x.UnwrapVal())
	}
}

func (e *Eval) fOr() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() | x.UnwrapVal())
	}
}

func (e *Eval) fXor() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(y.UnwrapVal() ^ x.UnwrapVal())
	}
}

func (e *Eval) invert() {
	x := e.Stack.Pop()
	if x.IsOk() {
		e.Stack.Push(^(x.UnwrapVal()))
	}
}

func (e *Eval) equal() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(If(y.UnwrapVal() == x.UnwrapVal()))
	}
}

func (e *Eval) notEqual() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(If(y.UnwrapVal() != x.UnwrapVal()))
	}
}

func (e *Eval) lessThan() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(If(y.UnwrapVal() < x.UnwrapVal()))
	}
}

func (e *Eval) greaterThan() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(If(y.UnwrapVal() > x.UnwrapVal()))
	}
}

func (e *Eval) lessThanEqual() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(If(y.UnwrapVal() <= x.UnwrapVal()))
	}
}

func (e *Eval) greaterThanEqual() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x.IsOk() && y.IsOk() {
		e.Stack.Push(If(y.UnwrapVal() >= x.UnwrapVal()))
	}
}

func (e *Eval) within() {
	upper := e.Stack.Pop()
	lower := e.Stack.Pop()
	val := e.Stack.Pop()

	if upper.IsOk() && lower.IsOk() && val.IsOk() {
		u := upper.UnwrapVal()
		l := lower.UnwrapVal()
		v := val.UnwrapVal()
		y := If(l < u && l <= v && v < u || l > u && (l <= v || v < u))
		e.Stack.Push(y)
	}
}
