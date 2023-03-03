package main

func (e *Eval) startDefinition() {
	e.compiling = true
}

func (e *Eval) endDefinition() {
	e.Dict[e.tmp.Name] = e.tmp
	e.tmp = Word{Name: ""}
	e.compiling = false
}
