package snaker

import (
	"fmt"
	"testing"
)

type I interface {
	Test()
}
type A struct {

}

func (a *A) Test() {
	fmt.Println("a")
}

type B struct {
	A
}

func (b *B) Test() {
	fmt.Println("b")
}

func TestAddTaskActor(t *testing.T) {
	i := new(B)
	i.Test()
}
