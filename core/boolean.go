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

func (e *Eval) ftrue() error {
	e.Stack.Push(TT)
	return nil
}

func (e *Eval) ffalse() error {
	e.Stack.Push(FF)
	return nil
}

func (e *Eval) fand() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	
	e.Stack.Push(y.UnwrapVal() & x.UnwrapVal())
	return nil
}

func (e *Eval) fOr() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(y.UnwrapVal() | x.UnwrapVal())
	return nil
}

func (e *Eval) fXor() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(y.UnwrapVal() ^ x.UnwrapVal())
	return nil
}

func (e *Eval) invert() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	e.Stack.Push(^(x.UnwrapVal()))
	return nil
}

func (e *Eval) equal() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(If(y.UnwrapVal() == x.UnwrapVal()))
	return nil
}

func (e *Eval) notEqual() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(If(y.UnwrapVal() != x.UnwrapVal()))
	return nil
}

func (e *Eval) lessThan() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(If(y.UnwrapVal() < x.UnwrapVal()))
	return nil
}

func (e *Eval) greaterThan() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(If(y.UnwrapVal() > x.UnwrapVal()))
	return nil
}

func (e *Eval) lessThanEqual() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(If(y.UnwrapVal() <= x.UnwrapVal()))
	return nil
}

func (e *Eval) greaterThanEqual() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(If(y.UnwrapVal() >= x.UnwrapVal()))
	return nil
}

func (e *Eval) within() error {
	upper := e.Stack.Pop()
	if !upper.IsOk() {
		return upper.UnwrapErr()
	}
	lower := e.Stack.Pop()
	if !lower.IsOk() {
		return lower.UnwrapErr()
	}
	val := e.Stack.Pop()
	if !val.IsOk() {
		return val.UnwrapErr()
	}

	u := upper.UnwrapVal()
	l := lower.UnwrapVal()
	v := val.UnwrapVal()
	y := If(l < u && l <= v && v < u || l > u && (l <= v || v < u))
	e.Stack.Push(y)
	return nil
}
