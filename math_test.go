package main

import (
	"bytes"
	"log"
	"testing"
)

var LogMsg = "Value %d given. Should be %d"

func expectedVal(t *testing.T, x Result[int64], expectedVal int64, log bytes.Buffer) {
	if !x.IsOk() {
		t.Fatalf(`Result not ok`)
	}

	if x.UnwrapVal() != expectedVal {
		t.Fatal(&log)
	}
}

func TestPlus(t *testing.T) {
	var a Result[int64]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestPlus: ", log.Lshortfile)
	midUint := int64(e.maxUint / 2)

	e.Stack.Push(0)
	e.Stack.Push(5)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 5)
	expectedVal(t, a, 5, buf)
	
	e.Stack.Push(5)
	e.Stack.Push(0)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 5)
	expectedVal(t, a, 5, buf) 
	
	e.Stack.Push(0)
	e.Stack.Push(-5)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -5)
	expectedVal(t, a, -5, buf) 
	
	e.Stack.Push(-5)
	e.Stack.Push(-0)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -5)
	expectedVal(t, a, -5, buf) 
	
	e.Stack.Push(1)
	e.Stack.Push(2)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 3)
	expectedVal(t, a, 3, buf) 
	
	e.Stack.Push(1)
	e.Stack.Push(-2)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf) 
	
	e.Stack.Push(-1)
	e.Stack.Push(2)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf) 
	
	e.Stack.Push(-1)
	e.Stack.Push(-2)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -3)
	expectedVal(t, a, -3, buf) 
	
	e.Stack.Push(-1)
	e.Stack.Push(1)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(),  0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(midUint)
	e.Stack.Push(1)
	e.plus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), midUint + 1)
	expectedVal(t, a, midUint + 1, buf)
}

func TestMinus(t *testing.T) {
	var a Result[int64]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestMinus: ", log.Lshortfile)
	midUint := int64(e.maxUint / 2)

	e.Stack.Push(0)
	e.Stack.Push(5)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -5)
	expectedVal(t, a, -5, buf)
	
	e.Stack.Push(5)
	e.Stack.Push(0)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 5)
	expectedVal(t, a, 5, buf) 
	
	e.Stack.Push(0)
	e.Stack.Push(-5)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 5)
	expectedVal(t, a, 5, buf) 
	
	e.Stack.Push(-5)
	e.Stack.Push(0)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -5)
	expectedVal(t, a, -5, buf) 
	
	e.Stack.Push(1)
	e.Stack.Push(2)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf) 
	
	e.Stack.Push(1)
	e.Stack.Push(-2)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 3)
	expectedVal(t, a, 3, buf) 
	
	e.Stack.Push(-1)
	e.Stack.Push(2)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -3)
	expectedVal(t, a, -3, buf) 
	
	e.Stack.Push(-1)
	e.Stack.Push(-2)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf) 
	
	e.Stack.Push(0)
	e.Stack.Push(1)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf)

	e.Stack.Push(midUint + 1)
	e.Stack.Push(1)
	e.minus()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), midUint)
	expectedVal(t, a, midUint, buf)
}

func TestStar(t *testing.T) {
	var a Result[int64]
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestStar: ", log.Lshortfile)
	midUint := int64(e.maxUint / 2)
	
	e.Stack.Push(0)
	e.Stack.Push(0)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(1)
	e.Stack.Push(0)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)
	
	e.Stack.Push(1)
	e.Stack.Push(2)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 2)
	expectedVal(t, a, 2, buf)
	
	e.Stack.Push(2)
	e.Stack.Push(1)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 2)
	expectedVal(t, a, 2, buf)
	
	e.Stack.Push(3)
	e.Stack.Push(3)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 9)
	expectedVal(t, a, 9, buf)
	
	e.Stack.Push(-3)
	e.Stack.Push(3)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -9)
	expectedVal(t, a, -9, buf)

	e.Stack.Push(3)
	e.Stack.Push(-3)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -9)
	expectedVal(t, a, -9, buf)
	
	e.Stack.Push(-3)
	e.Stack.Push(-3)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 9)
	expectedVal(t, a, 9, buf)

	e.Stack.Push(midUint + 1)
	e.Stack.Push(1)
	e.rShift()
	e.Stack.Push(2)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), midUint + 1)
	expectedVal(t, a, midUint + 1, buf)

	e.Stack.Push(midUint + 1)
	e.Stack.Push(2)
	e.rShift()
	e.Stack.Push(4)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), midUint + 1)
	expectedVal(t, a, midUint + 1, buf)

	e.Stack.Push(midUint +1)
	e.Stack.Push(1)
	e.rShift()
	e.Stack.Push(midUint + 1)
	e.fOr()
	e.Stack.Push(2)
	e.star()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), midUint + 1)
	expectedVal(t, a, midUint + 1, buf)
}

func TestSlash(t *testing.T) {
	var a Result[int64]
	var buf bytes.Buffer
	e := NewEval()
	logger := log.New(&buf, "TestSlash: ", log.Lshortfile)

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(1)
	e.Stack.Push(1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(2)
	e.Stack.Push(1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 2)
	expectedVal(t, a, 2, buf)

	e.Stack.Push(-1)
	e.Stack.Push(1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf)

	e.Stack.Push(-2)
	e.Stack.Push(1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -2)
	expectedVal(t, a, -2, buf)

	e.Stack.Push(0)
	e.Stack.Push(-1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(1)
	e.Stack.Push(-1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf)

	e.Stack.Push(2)
	e.Stack.Push(-1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -2)
	expectedVal(t, a, -2, buf)

	e.Stack.Push(-1)
	e.Stack.Push(-1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(-2)
	e.Stack.Push(-1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 2)
	expectedVal(t, a, 2, buf)

	e.Stack.Push(2)
	e.Stack.Push(2)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(2)
	e.Stack.Push(2)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(7)
	e.Stack.Push(3)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 2)
	expectedVal(t, a, 2, buf)

	e.Stack.Push(e.maxInt)
	e.Stack.Push(1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), e.maxInt)
	expectedVal(t, a, e.maxInt, buf)

	e.Stack.Push(e.minInt)
	e.Stack.Push(1)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), e.minInt)
	expectedVal(t, a, e.minInt, buf)

	e.Stack.Push(e.maxInt)
	e.Stack.Push(e.maxInt)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(e.minInt)
	e.Stack.Push(e.minInt)
	e.slash()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)
}

func TestMod(t *testing.T) {
	var a Result[int64]
	var buf bytes.Buffer
	e := NewEval()
	logger := log.New(&buf, "TestMod: ", log.Lshortfile)

	e.Stack.Push(0)
	e.Stack.Push(1)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(1)
	e.Stack.Push(1)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(2)
	e.Stack.Push(1)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)


	e.Stack.Push(-2)
	e.Stack.Push(1)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(7)
	e.Stack.Push(3)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(7)
	e.Stack.Push(-3)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 1)
	expectedVal(t, a, 1, buf)

	e.Stack.Push(-7)
	e.Stack.Push(3)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf)

	e.Stack.Push(-7)
	e.Stack.Push(-3)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), -1)
	expectedVal(t, a, -1, buf)

	e.Stack.Push(e.maxInt)
	e.Stack.Push(1)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(e.minInt)
	e.Stack.Push(1)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(e.maxInt)
	e.Stack.Push(e.maxInt)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(e.minInt)
	e.Stack.Push(e.minInt)
	e.mod()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)
}

func TestTwoStar(t *testing.T) {
	var a Result[int64]
	var buf bytes.Buffer
	e := NewEval()
	logger := log.New(&buf, "TestTwoStar: ", log.Lshortfile)

	e.Stack.Push(0)
	e.twoStar()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 0)
	expectedVal(t, a, 0, buf)

	e.Stack.Push(1)
	e.twoStar()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 2)
	expectedVal(t, a, 2, buf)

	e.Stack.Push(4000)
	e.twoStar()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 8000)
	expectedVal(t, a, 8000, buf)

	e.Stack.Push(1)
	e.twoStar()
	e.Stack.Push(1)
	e.fXor()
	a = e.Stack.Pop()
	logger.Printf(LogMsg, a.UnwrapVal(), 3)
	expectedVal(t, a, 3, buf)
}
