package core

import (
	"fmt"
	"math"
	"strings"
	"strconv"

	"github.com/MorganPeterson/mForth/stack"
)

type Word struct {
	Name string
	Func func() error
	Words []string
}

type Eval struct {
	FStack stack.Stack[float32]
	Stack stack.Stack[int]
	RStack stack.Stack[int]
	Dict map[string]Word
	compiling bool
	tmp Word
	maxUint uint
	maxInt int
	minInt int
	comment int
}

func NewEval() *Eval {
	e := &Eval{
		tmp: Word{Name: ""},
		maxUint: math.MaxUint,
		maxInt: math.MaxInt,
		minInt: math.MinInt,
	}
	
	e.Dict = map[string]Word{
		"+": {Name: "+", Func: e.plus},
		"1+": {Name: "1+", Func: e.onePlus},
		"-": {Name: "-", Func: e.minus},
		"1-": {Name: "1-", Func: e.oneMinus},
		"*": {Name: "*", Func: e.star},
		"/": {Name: "/", Func: e.slash},
		".": {Name: ".", Func: e.dot},
		":": {Name: ":", Func: e.startDefinition},
		";": {Name: ";", Func: e.endDefinition},
		"DROP": {Name: "DROP", Func: e.drop},
		"2DROP": {Name: "2DROP", Func: e.twoDrop},
		"SWAP": {Name: "SWAP", Func: e.swap},
		"2SWAP": {Name: "2SWAP", Func: e.twoSwap},
		"DUP": {Name: "DUP", Func: e.dup},
		"2DUP": {Name: "2DUP", Func: e.twoDup},
		"?DUP": {Name: "?DUP", Func: e.questionDup},
		"OVER": {Name: "OVER", Func: e.over},
		"2OVER": {Name: "2OVER", Func: e.twoOver},
		"PICK": {Name: "PICK", Func: e.pick},
		"ROT": {Name: "ROT", Func: e.rot},
		"ROLL": {Name: "ROLL", Func: e.roll},
		"DEPTH": {Name: "DEPTH", Func: e.depth},
		">R": {Name: ">R", Func: e.toR},
		"R>": {Name: "R>", Func: e.fromR},
		"R@": {Name: "R@", Func: e.fetchR},
		"2R>": {Name: "2R>", Func: e.twoFromR},
		"2>R": {Name: "2>R", Func: e.twoToR},
		"@": {Name: "@", Func: e.fetch},
		"2@": {Name: "2@", Func: e.twoFetch},
		"2R@": {Name: "2R@", Func: e.fetchTwoR},
		"TRUE": {Name: "TRUE", Func: e.ftrue},
		"FALSE": {Name: "FALSE", Func: e.ffalse},
		"AND": {Name: "AND", Func: e.fand},
		"OR": {Name: "OR", Func: e.fOr},
		"XOR": {Name: "XOR", Func: e.fXor},
		"INVERT": {Name: "INVERT", Func: e.invert},
		"=": {Name: "=", Func: e.equal},
		"<": {Name: "<", Func: e.lessThan},
		">": {Name: ">", Func: e.greaterThan},
		"<>": {Name: "<>", Func: e.notEqual},
		"WITHIN": {Name: "WITHIN", Func: e.within},
		"0<": {Name: "0<", Func: e.zeroLess},
		"0=": {Name: "0=", Func: e.zeroEquals},
		"2*": {Name: "2*", Func: e.twoStar},
		"2/": {Name: "2/", Func: e.twoSlash},
		"MOD": {Name: "MOD", Func: e.mod},
		"*/MOD": {Name: "*/MOD", Func: e.starSlashMod},
		"!": {Name: "!", Func: e.store},
		"RSHIFT": {Name: "RSHIFT", Func: e.rShift},
		"(": {Name: "(", Func: e.leftParen},
	}
	return e
}

func (e *Eval) Eval(args []string) {
	for _, tok := range args {
		tok = strings.TrimSpace(tok)
		
		if tok == "" {
			continue
		}

		if tok == "\\" {
			break
		}

		if e.comment > 0 {
			if tok == ")" {
				e.comment--
			}
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
					_, err := strconv.ParseInt(tok, 0, 64)
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
			i, err := strconv.Atoi(tok)
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
	var fResult error
	
	if word.Func != nil {
		fResult = word.Func()
		if fResult != nil  {
			fmt.Printf("Eval: %s: %s\n", fResult.Error(), word.Name)
		}
	} else {
		addNum := false
		for _, offset := range word.Words {
			if addNum {
				n, err := strconv.Atoi(offset)
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
						fmt.Printf("%s: %s not found\n", offset, offset)
						break
					}
				}
			}
		}
	}
}
