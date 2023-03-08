package core

import (
	"bytes"
	"log"
	"testing"

	"github.com/MorganPeterson/mForth/result"
)

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
