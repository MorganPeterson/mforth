package main

import "math"

var (
	TT = float64(-1)
	FF = float64(0)
)

func If(cond bool) float64 {
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

	if x == nil || y == nil {
		return
	}
	xx := int64(math.Float64bits(*x))
	yy := int64(math.Float64bits(*y))

	z := yy & xx
	e.Stack.Push(math.Float64frombits(uint64(z)))
}

func (e *Eval) fOr() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x == nil || y == nil {
		return
	}
	xx := int64(math.Float64bits(*x))
	yy := int64(math.Float64bits(*y))

	z := yy | xx
	e.Stack.Push(math.Float64frombits(uint64(z)))
}

func (e *Eval) fXor() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()

	if x == nil || y == nil {
		return
	}
	xx := int64(math.Float64bits(*x))
	yy := int64(math.Float64bits(*y))

	z := yy ^ xx
	e.Stack.Push(math.Float64frombits(uint64(z)))
}

func (e *Eval) invert() {
	x := e.Stack.Pop()
	if x == nil {
		return
	}
	y := ^(int64(math.Float64bits(*x)))
	e.Stack.Push(math.Float64frombits(uint64(y)))
}

func (e *Eval) equal() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x == nil || y == nil {
		return
	}
	e.Stack.Push(If(*y == *x))
}

func (e *Eval) notEqual() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x == nil || y == nil {
		return
	}
	e.Stack.Push(If(*y != *x))
}

func (e *Eval) lessThan() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x == nil || y == nil {
		return
	}
	e.Stack.Push(If(*y < *x))
}

func (e *Eval) greaterThan() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x == nil || y == nil {
		e.Stack.Push(If(*y > *x))
	}
}

func (e *Eval) lessThanEqual() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x == nil || y == nil {
		return
	}
	e.Stack.Push(If(*y <= *x))
}

func (e *Eval) greaterThanEqual() {
	x := e.Stack.Pop()
	y := e.Stack.Pop()
	if x != nil && y != nil {
		e.Stack.Push(If(*y >= *x))
	}
}

func (e *Eval) within() {
	upper := e.Stack.Pop()
	lower := e.Stack.Pop()
	val := e.Stack.Pop()
	if upper == nil && lower == nil && val == nil {
		return
	}
	u := *upper
	l := *lower
	v := *val
	y := If(l < u && l <= v && v < u || l > u && (l <= v || v < u))
	e.Stack.Push(y)
}
