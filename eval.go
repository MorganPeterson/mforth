package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Word struct {
	Name string
	Func func()
	Words []string
}

type Eval struct {
	Stack Stack[float64]
	RStack Stack[float64]
	Dict map[string]Word
	compiling bool
	tmp Word
}

func NewEval() *Eval {
	e := &Eval{tmp: Word{Name: ""},}
	e.Dict = map[string]Word{
		"+": {Name: "+", Func: e.add},
		"-": {Name: "-", Func: e.sub},
		"*": {Name: "*", Func: e.mul},
		"/": {Name: "/", Func: e.div},
		".": {Name: ".", Func: e.print},
		":": {Name: ":", Func: e.startDefinition},
		";": {Name: ";", Func: e.endDefinition},
		"PRINT": {Name: "PRINT", e.print},
		"DROP": {Name: "DROP", Func: e.drop},
		"2DROP": {Name: "2DROP", Func: e.twoDrop},
		"SWAP": {Name: "SWAP", Func: e.swap},
		"2SWAP": {Name: "2SWAP", Func: e.twoSwap},
		"DUP": {Name: "DUP", Func: e.dup},
		"2DUP": {Name: "2DUP", Func: e.twoDup},
		"?DUP": {Name: "?DUP", Func: e.nonZeroDup},
		"OVER": {Name: "OVER", Func: e.over},
		"2OVER": {Name: "2OVER", Func: e.twoOver},
		"PICK": {Name: "PICK", Func: e.pick},
		"ROT": {Name: "ROT", Func: e.rot},
		"-ROT": {Name: "-ROT", Func: e.reverseRot},
		"ROLL": {Name: "ROLL", Func: e.roll},
		"DEPTH": {Name: "DEPTH", Func: e.depth},
		">R": {Name: ">R", Func: e.toR},
		"R>": {Name: "R>", Func: e.fromR},
		"R@": {Name: "R@", Func: e.fetchR},
		"2R>": {Name: "2R>", Func: e.twoFromR},
		"2>R": {Name: "2>R", Func: e.twoToR},
		"2R@": {Name: "2R@", Func: e.fetchTwoR},
	}
	return e
}

func (e *Eval) Eval(args []string) {
	for _, tok := range args {
		tok = strings.TrimSpace(tok)
		if tok == "" {
			continue
		}

		if e.compiling {
			if tok == ";" {
				word, _ := e.findWord(tok)
				e.evalWord(word)
				continue
			} else {
				if e.tmp.Name == "" {
					_, prs := e.findWord(tok)
					if prs {
						fmt.Printf("word %s already defined\n", tok)
						return
					}
					e.tmp.Name = tok
					continue
				}
				
				val, prs := e.findWord(tok)
				if prs {
					e.tmp.Words = append(e.tmp.Words, val.Name)
				} else {
					e.tmp.Words = append(e.tmp.Words, "")
					_, err := strconv.ParseFloat(tok, 64)
					if err != nil {
						fmt.Printf("%s: %s\n", tok, err.Error())
						return
					}
					e.tmp.Words = append(e.tmp.Words, tok)
				}
				continue
			}
		}

		wrd, prs := e.findWord(tok)
		if prs {
			e.evalWord(wrd)
		} else {
			i, err := strconv.ParseFloat(tok, 64)
			if err != nil {
				fmt.Printf("%s: %s\n", tok, err.Error())
				return
			}
			e.Stack.Push(i)
		}
	}
}

func (e *Eval) findWord(name string) (Word, bool) {
	f, prs := e.Dict[name]
	return f, prs
}

func (e *Eval) evalWord(word Word) {
	if word.Func != nil {
		word.Func()
	} else {
		addNum := false
		for _, offset := range word.Words {
			if addNum {
				n, err := strconv.ParseFloat(offset, 64)
				if err == nil {
					e.Stack.Push(n)
					addNum = false
				}
			} else {
				if offset == "" {
					addNum = true
				} else {
					f, p := e.findWord(offset)
					if p {
						e.evalWord(f)
					} else {
						fmt.Printf("%s: %s not found\n", f)
						break
					}
				}
			}
		}
	}
}

