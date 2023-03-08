package core

import (
	"errors"
	"sync"

	"github.com/MorganPeterson/mForth/result"
)

func checkIsOk[T any](wg *sync.WaitGroup, err chan error, x *result.Result[T]) {
	defer wg.Done()
	if !x.IsOk() {
		err <- x.UnwrapErr()
	} else {
		err <- nil
	}
}

func (e *Eval) plus() error {
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

	e.Stack.Push(y.UnwrapVal() + x.UnwrapVal())
	return nil
}

func (e *Eval) minus() error {
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

	e.Stack.Push(y.UnwrapVal() - x.UnwrapVal())
	return nil
}

func (e *Eval) star() error {
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

	e.Stack.Push(y.UnwrapVal() * x.UnwrapVal())
	return nil
}

func (e *Eval) slash() error {
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
	if xx == 0 {
		return errors.New("SLASH: zero division error")
	}
	e.Stack.Push(y.UnwrapVal() / xx)
	return nil
}

func (e *Eval) mod() error {
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

	xx := x.UnwrapVal()
	if xx == 0 {
		return errors.New("STARSLASHMOD: zero division error")
	}
	a := z.UnwrapVal() * y.UnwrapVal()
	e.Stack.Push(a % xx)
	e.Stack.Push(a / xx)

	return nil
}

func (e *Eval) onePlus() error {
	x := e.Stack.Pop()

	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal() + 1)
	}

	return x.UnwrapErr()
}

func (e *Eval) oneMinus() error {
	x := e.Stack.Pop()

	if x.IsOk() {
		e.Stack.Push(x.UnwrapVal() - 1)
	}

	return x.UnwrapErr()
}
