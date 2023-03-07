package core

import (
	"errors"
	"fmt"
)

func (e *Eval) startDefinition() error {
	e.compiling = true
	return nil
}

func (e *Eval) endDefinition() error {
	e.Dict[e.tmp.Name] = e.tmp
	e.tmp = Word{Name: ""}
	e.compiling = false
	return nil
}


func (e *Eval) dot() error {
	x := e.Stack.Pop()
	if x.IsOk() {
		fmt.Printf("%d\n", x.UnwrapVal())
		return nil
	}
	return x.UnwrapErr()
}

func (e *Eval) drop() error {
	r := e.Stack.Pop()
	if r.IsOk() {
		return nil
	}
	return r.UnwrapErr()
}

func (e *Eval) twoDrop() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	return nil
}

func (e *Eval) swap() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(x.UnwrapVal())
	e.Stack.Push(y.UnwrapVal())
	return nil
}

func (e *Eval) twoSwap() error {
	v := e.Stack.Pop()
	if !v.IsOk() {
		return v.UnwrapErr()
	}

	w := e.Stack.Pop()
	if !w.IsOk() {
		return w.UnwrapErr()
	}

	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(w.UnwrapVal())
	e.Stack.Push(v.UnwrapVal())
	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())

	return nil
}

func (e *Eval) dup() error {
	x := e.Stack.Peek()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	e.Stack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) twoDup() error {
	x := e.Stack.Peek()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	y := e.Stack.Fetch(2)
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())

	return nil
}

func (e *Eval) nonZeroDup() error {
	x := e.Stack.Peek()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	xx := x.UnwrapVal()
	if xx != 0 {
		e.Stack.Push(xx)
	}
	return nil
}

func (e *Eval) over() error {
	x := e.Stack.Fetch(e.Stack.Len() - 2)
	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal())
		return nil
	}
	return x.UnwrapErr()
}

func (e *Eval) twoOver() error {
	x := e.Stack.Fetch(3)
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	y := e.Stack.Fetch(4)
	if !y.IsOk() {
		return x.UnwrapErr()
	}
	
	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) pick() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	xx := x.UnwrapVal() + 1
	y := e.Stack.Fetch(e.Stack.Len() - xx)
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	e.Stack.Push(y.UnwrapVal())
	return nil
}

func (e *Eval) rot() error {
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

	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())
	e.Stack.Push(z.UnwrapVal())
	return nil
}

func (e *Eval) reverseRot() error {
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

	e.Stack.Push(x.UnwrapVal())
	e.Stack.Push(z.UnwrapVal())
	e.Stack.Push(y.UnwrapVal())
	return nil
}

func (e *Eval) depth() error {
	e.Stack.Push(e.Stack.Len())
	return nil
}

func (e *Eval) roll() error {
	x := e.Stack.Pop()

	if !x.IsOk() {
		return x.UnwrapErr()
	}

	xx := x.UnwrapVal()

	switch xx {
	case 0:
		return nil
	case 1:
		return e.swap()
	case 2:
		return e.rot()
	default:
		i := e.Stack.Len() - (xx + 1)
		y := e.Stack.PeekAt(i)
		e.Stack.Rm(i)
		e.Stack.Insert(e.Stack.Len()-1, y)
	}
	return nil
}

/*
 * return stack functions
 */

func (e *Eval) toR() error {
	x := e.Stack.Pop()
	
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	e.RStack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) fromR() error {
	x := e.RStack.Pop()

	if !x.IsOk() {
		return x.UnwrapErr()
	}

	e.Stack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) fetchR() error {
	x := e.RStack.Peek()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	e.Stack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) twoToR() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	e.RStack.Push(y.UnwrapVal())
	e.RStack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) twoFromR() error {
	x := e.RStack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}
	
	y := e.RStack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) fetchTwoR() error {
	x := e.RStack.Fetch(2)
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	y := e.RStack.Peek()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	e.Stack.Push(x.UnwrapVal())
	e.Stack.Push(y.UnwrapVal())
	return nil
}

func (e *Eval) leftParen() error {
	e.comment++
	return nil
}

func (e *Eval) store() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}

	xx := x.UnwrapVal()
	if xx < 0 {
		errors.New("!: STORE: Illegal attempt to change input")
	}

	e.Stack.Insert(xx, y.UnwrapVal())
	return nil
}

func (e *Eval) rShift() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	y := e.Stack.Pop()
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	e.Stack.Push(y.UnwrapVal() >> x.UnwrapVal())
	return nil
}
