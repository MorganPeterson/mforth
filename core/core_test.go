package core

import (
	"bytes"
	"log"
	"testing"
)

func TestDrop(t *testing.T) {
	var a int
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestDrop: ", log.Lshortfile)

	e.Stack.Push(1)
	e.Stack.Push(2)
	e.drop()
	a = e.Stack.Len()
	
	logger.Printf(LogMsg, a, 1)
	if a != 1 {
		t.Fatal(&buf)
	}

	e.Stack.Push(0)
	e.drop()
	a = e.Stack.Len()

	logger.Printf(LogMsg, a, 1)
	if a != 1 {
		t.Fatal(&buf)
	}
}

func TestTwoDrop(t *testing.T) {
	var a int
	var buf bytes.Buffer

	e := NewEval()
	logger := log.New(&buf, "TestTwoDrop: ", log.Lshortfile)

	e.Stack.Push(1)
	e.Stack.Push(2)
	e.twoDrop()
	a = e.Stack.Len()
	
	logger.Printf(LogMsg, a, 0)
	if a != 0 {
		t.Fatal(&buf)
	}
}
