package core

import (
	"errors"
)

func (e *Eval) plus() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(y.UnwrapVal() + x.UnwrapVal())
	return nil
}

func (e *Eval) minus() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(y.UnwrapVal() - x.UnwrapVal())
	return nil
}

func (e *Eval) star() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(y.UnwrapVal() * x.UnwrapVal())
	return nil
}

func (e *Eval) slash() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	xx := x.UnwrapVal()
	if xx == 0 {
		return errors.New("SLASH: zero division error")
	}
	e.Stack.Push(y.UnwrapVal() / xx)
	return nil
}

func (e *Eval) mod() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	xx := x.UnwrapVal()
	if xx == 0 {
		return errors.New("SLASH: zero division error")
	}

	e.Stack.Push(y.UnwrapVal() % xx)
	return nil
}

func (e *Eval) twoStar() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	e.Stack.Push(x.UnwrapVal() << 1)
	return nil
}

func (e *Eval) twoSlash() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	e.Stack.Push(x.UnwrapVal() >> 1)
	return nil
}

func (e *Eval) starSlashMod() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	z := e.Stack.Pop()
	if !z.IsOk() {
		return z.UnwrapErr()
	}

	xx := x.UnwrapVal()
	if xx == 0 {
		return errors.New("STARSLASHMOD: zero division error")
	}
	a := z.UnwrapVal() * y.UnwrapVal()
	e.Stack.Push(a % xx)
	e.Stack.Push(a / xx)

	return nil
}

func (e *Eval) plusStore() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	
	e.Stack.Push(y.UnwrapVal() + x.UnwrapVal())
	return nil
}
		
