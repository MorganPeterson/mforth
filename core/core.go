package core

import (
	"errors"
	"fmt"
	"sync"
)

// colon starts a word definition
func (e *Eval) startDefinition() error {
	e.compiling = true
	return nil
}

// semicolon ends a word definition
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
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	return isErr
}

func (e *Eval) swap() error {
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	e.Stack.Push(x.UnwrapVal())
	e.Stack.Push(y.UnwrapVal())
	return nil
}

func (e *Eval) twoSwap() error {
	err := make(chan error, 4)
	done := sync.WaitGroup{}

	v := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &v)

	w := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &w)

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
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
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())

	return nil
}

func (e *Eval) questionDup() error {
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
	x := e.Stack.PeekAt(e.Stack.Len() - 2)
	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal())
		return nil
	}
	return x.UnwrapErr()
}

func (e *Eval) twoOver() error {
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
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
	y := e.Stack.PeekAt(e.Stack.Len() - xx)
	if !y.IsOk() {
		return y.UnwrapErr()
	}
	e.Stack.Push(y.UnwrapVal())
	return nil
}

func (e *Eval) rot() error {
	err := make(chan error, 3)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	z := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &z)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())
	e.Stack.Push(z.UnwrapVal())
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
		if !y.IsOk() {
			return y.UnwrapErr()
		}
		e.Stack.Rm(i)
		e.Stack.Insert(e.Stack.Len()-1, y.UnwrapVal())
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
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	e.RStack.Push(y.UnwrapVal())
	e.RStack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) twoFromR() error {
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.RStack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.RStack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	e.Stack.Push(y.UnwrapVal())
	e.Stack.Push(x.UnwrapVal())
	return nil
}

func (e *Eval) fetchTwoR() error {
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.RStack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.RStack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
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
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	xx := x.UnwrapVal()
	if xx < 0 {
		errors.New("!: STORE: Illegal attempt to change input")
	}

	e.Stack.Insert(xx, y.UnwrapVal())
	return nil
}

func (e *Eval) rShift() error {
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	x := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &x)

	y := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &y)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	e.Stack.Push(y.UnwrapVal() >> x.UnwrapVal())
	return nil
}

func (e *Eval) twoFetch() error {
	err := make(chan error, 2)
	done := sync.WaitGroup{}

	addr := e.Stack.Pop()
	if !addr.IsOk() {
		return addr.UnwrapErr()
	}

	addrVal := addr.UnwrapVal()

	x1 := e.Stack.PeekAt(addrVal)
	done.Add(1)
	go checkIsOk[int](&done, err, &x1)

	x2 := e.Stack.PeekAt(addrVal + 1)
	done.Add(1)
	go checkIsOk[int](&done, err, &x2)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	fmt.Printf("%d\n", addrVal)
	fmt.Printf("x1: %d -> %d, x2: %d -> %d\n", x1.UnwrapVal(), addrVal + 1, x2.UnwrapVal(), addrVal)

	e.Stack.Insert(addrVal, x2.UnwrapVal())
	e.Stack.Insert(addrVal + 1, x1.UnwrapVal())

	return nil
}

func (e *Eval) fetch() error {
	addr := e.Stack.Pop()
	if !addr.IsOk() {
		return addr.UnwrapErr()
	}

	addrVal := addr.UnwrapVal()

	x1 := e.Stack.PeekAt(addrVal)
	if !x1.IsOk() {
		return x1.UnwrapErr()
	}

	e.Stack.Push(x1.UnwrapVal())
	return nil
}
