package core

import "sync"

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

	e.Stack.Push(y.UnwrapVal() & x.UnwrapVal())
	return nil
}

func (e *Eval) fOr() error {
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

	e.Stack.Push(y.UnwrapVal() | x.UnwrapVal())
	return nil
}

func (e *Eval) fXor() error {
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

	e.Stack.Push(y.UnwrapVal() ^ x.UnwrapVal())
	return nil
}

func (e *Eval) invert() error {
	x := e.Stack.Pop()
	if !x.IsOk() {
		return x.UnwrapErr()
	}

	e.Stack.Push(^x.UnwrapVal())
	return nil
}

func (e *Eval) equal() error {
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

	e.Stack.Push(If(y.UnwrapVal() == x.UnwrapVal()))
	return nil
}

func (e *Eval) notEqual() error {
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

	e.Stack.Push(If(y.UnwrapVal() != x.UnwrapVal()))
	return nil
}

func (e *Eval) lessThan() error {
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

	e.Stack.Push(If(y.UnwrapVal() < x.UnwrapVal()))
	return nil
}

func (e *Eval) greaterThan() error {
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

	e.Stack.Push(If(y.UnwrapVal() > x.UnwrapVal()))
	return nil
}

func (e *Eval) within() error {
	err := make(chan error, 3)
	done := sync.WaitGroup{}

	upper := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &upper)

	lower := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &lower)

	val := e.Stack.Pop()
	done.Add(1)
	go checkIsOk[int](&done, err, &val)

	done.Wait()

	isErr := <-err
	close(err)

	if isErr != nil {
		return isErr
	}

	u := upper.UnwrapVal()
	l := lower.UnwrapVal()
	v := val.UnwrapVal()
	y := If(l < u && l <= v && v < u || l > u && (l <= v || v < u))
	e.Stack.Push(y)
	return nil
}

func (e *Eval) zeroLess() error {
	x := e.Stack.Pop()
	if x.IsOk() {
		e.Stack.Push(If(x.UnwrapVal() < 0))
		return nil
	}
	return x.UnwrapErr()
}

func (e *Eval) zeroEquals() error {
	x := e.Stack.Pop()
	if x.IsOk() {
		e.Stack.Push(If(x.UnwrapVal() == 0))
		return nil
	}
	return x.UnwrapErr()
}
