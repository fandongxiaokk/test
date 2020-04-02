package interfacepolymorphic

import (
	"fmt"
	"testing"
)

type Code string
type interA interface {
	PromentA(host Code) Code
}

type PromentAB struct {
}

func (p *PromentAB) PromentA(host2 Code) Code {
	return "this is PromentAB"
}

type PromentABC struct {
}

func (p *PromentABC) PromentA(host1 Code) Code {
	// return byes.join("this is ", host)
	return host1
}

func writeFirstProgram(w interA, t Code) {
	fmt.Printf("%T %v \n", w, w.PromentA(t))
}

func TestA(t *testing.T) {
	k := &PromentABC{}
	u := new(PromentAB)
	writeFirstProgram(u, "")
	writeFirstProgram(k, "fan")
}
