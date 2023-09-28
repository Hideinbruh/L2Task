package main

import "fmt"

type Visitor interface {
	VisitElementA(element *ElementA)
	VisitElementB(element *ElementB)
}

type Element interface {
	Accept(visitor Visitor)
}

type ElementA struct{}

func (e *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

func (e *ElementA) OperationA() {
	fmt.Println("Operation A in Element A")
}

type ElementB struct{}

func (e *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

func (e *ElementB) OperationB() {
	fmt.Println("Operation B in Element B")
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitElementA(element *ElementA) {
	element.OperationA()
}

func (v *ConcreteVisitor) VisitElementB(element *ElementB) {
	element.OperationB()
}

func main() {
	elements := []Element{&ElementA{}, &ElementB{}}
	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}
