package main

import "fmt"

type SubSystemA struct {
}

func (s *SubSystemA) OperationA() string {
	return "OperationA running\n"
}

type SubSystemB struct {
}

func (s *SubSystemB) OperationB() string {
	return "OperationB running\n"
}

type Facade struct {
	systemA *SubSystemA
	systemB *SubSystemB
}

func NewFacade() *Facade {
	return &Facade{
		systemA: &SubSystemA{},
		systemB: &SubSystemB{},
	}
}

func (f *Facade) Operation() string {
	result := "Facade initialized operations\n"
	result += f.systemA.OperationA()
	result += f.systemB.OperationB()
	return result
}

func main() {
	facade := NewFacade()
	result := facade.Operation()
	fmt.Println(result)
}
