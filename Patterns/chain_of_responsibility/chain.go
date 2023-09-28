package main

import "fmt"

// Интерфейс обработчика запросов
type Handler interface {
	handle(request string)
	setNext(handler Handler)
}

// Обработчик запросов "А"
type HandlerA struct {
	next Handler
}

func (h *HandlerA) handle(request string) {
	if request == "A" {
		fmt.Println("Обработчик А обработал запрос")
	} else if h.next != nil {
		h.next.handle(request)
	}
}

func (h *HandlerA) setNext(handler Handler) {
	h.next = handler
}

// Обработчик запросов "B"
type HandlerB struct {
	next Handler
}

func (h *HandlerB) handle(request string) {
	if request == "B" {
		fmt.Println("Обработчик B обработал запрос")
	} else if h.next != nil {
		h.next.handle(request)
	}
}

func (h *HandlerB) setNext(handler Handler) {
	h.next = handler
}

// Обработчик запросов "C"
type HandlerC struct {
	next Handler
}

func (h *HandlerC) handle(request string) {
	if request == "C" {
		fmt.Println("Обработчик C обработал запрос")
	} else if h.next != nil {
		h.next.handle(request)
	}
}

func (h *HandlerC) setNext(handler Handler) {
	h.next = handler
}

func main() {
	handlerA := &HandlerA{}
	handlerB := &HandlerB{}
	handlerC := &HandlerC{}

	handlerA.setNext(handlerB)
	handlerB.setNext(handlerC)

	handlerA.handle("A") // Обработчик А обработал запрос
	handlerA.handle("B") // Обработчик B обработал запрос
	handlerA.handle("C") // Обработчик C обработал запрос
	handlerA.handle("D") // Запрос не был обработан
}
