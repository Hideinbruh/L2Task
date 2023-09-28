package main

import "fmt"

// Интерфейс стратегии
type Strategy interface {
	doOperation(int, int) int
}

// Стратегия "Сложение"
type AddStrategy struct{}

func (s *AddStrategy) doOperation(num1 int, num2 int) int {
	return num1 + num2
}

// Стратегия "Вычитание"
type SubtractStrategy struct{}

func (s *SubtractStrategy) doOperation(num1 int, num2 int) int {
	return num1 - num2
}

// Контекст калькулятора
type Calculator struct {
	strategy Strategy
}

func (c *Calculator) setStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Calculator) executeStrategy(num1 int, num2 int) int {
	return c.strategy.doOperation(num1, num2)
}

func main() {
	calculator := &Calculator{}
	addStrategy := &AddStrategy{}
	subtractStrategy := &SubtractStrategy{}

	calculator.setStrategy(addStrategy)
	fmt.Println("10 + 5 =", calculator.executeStrategy(10, 5)) // 10 + 5 = 15

	calculator.setStrategy(subtractStrategy)
	fmt.Println("10 - 5 =", calculator.executeStrategy(10, 5)) // 10 - 5 = 5
}
