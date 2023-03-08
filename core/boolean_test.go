package core

import (
	"bytes"
	"log"
	"testing"

	"github.com/MorganPeterson/mForth/result"
)

func TestTrue(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestTrue: ", log.Lshortfile)

	e.ftrue()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), TT)
	expectedVal(t, a, TT, buf)

	e.ftrue()
	e.invert()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), FF)
	expectedVal(t, a, FF, buf)
}
	

func TestFalse(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestFalse: ", log.Lshortfile)

	e.ffalse()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), FF)
	expectedVal(t, a, FF, buf)
}

func TestFAnd(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestFand: ", log.Lshortfile)

	e.Stack.Push(0)
	e.Stack.Push(0)
	e.fand()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.fand()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(1)
	e.Stack.Push(0)
	e.fand()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(1)
	e.Stack.Push(1)
	e.fand()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)
}


func TestFOr(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestFOr: ", log.Lshortfile)

	e.Stack.Push(0)
	e.Stack.Push(0)
	e.fOr()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.fOr()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(1)
	e.Stack.Push(0)
	e.fOr()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(1)
	e.Stack.Push(1)
	e.fOr()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)
}

func TestFXor(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestFXor: ", log.Lshortfile)

	e.Stack.Push(0)
	e.Stack.Push(0)
	e.fXor()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.fXor()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(1)
	e.Stack.Push(0)
	e.fXor()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(1)
	e.Stack.Push(1)
	e.fXor()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)
}

func TestInvert(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestInvert: ", log.Lshortfile)

	e.Stack.Push(0)
	e.invert()

	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf)

	e.Stack.Push(1)
	e.invert()

	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -2)
	expectedVal(t, a, -2, buf)
}

func TestEquals(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestEqual: ", log.Lshortfile)

	e.Stack.Push(0)
	e.Stack.Push(0)
	e.equal()

	e.Stack.Push(1)
	e.Stack.Push(1)
	e.equal()

	e.Stack.Push(-1)
	e.Stack.Push(-1)
	e.equal()

	e.Stack.Push(1)
	e.Stack.Push(0)
	e.equal()

	e.Stack.Push(-1)
	e.Stack.Push(0)
	e.equal()

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.equal()

	e.Stack.Push(0)
	e.Stack.Push(-1)
	e.equal()

	x := []int{FF, FF, FF, FF, TT, TT, TT}

	for _, v := range x {
		a = e.Stack.Pop()
		logger.Printf(LogMsg, a.UnwrapVal(), v)
		expectedVal(t, a, v, buf)
	}
}

func TestNotEqual(t *testing.T) {
	var a result.Result[int]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestNotEqual: ", log.Lshortfile)

	e.Stack.Push(0)
	e.Stack.Push(0)
	e.notEqual()

	e.Stack.Push(1)
	e.Stack.Push(1)
	e.notEqual()

	e.Stack.Push(-1)
	e.Stack.Push(-1)
	e.notEqual()

	e.Stack.Push(1)
	e.Stack.Push(0)
	e.notEqual()

	e.Stack.Push(-1)
	e.Stack.Push(0)
	e.notEqual()

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.notEqual()

	e.Stack.Push(0)
	e.Stack.Push(-1)
	e.notEqual()

	x := []int{TT, TT, TT, TT, FF, FF, FF}

	for _, v := range x {
		a = e.Stack.Pop()
		logger.Printf(LogMsg, a.UnwrapVal(), v)
		expectedVal(t, a, v, buf)
	}
}
