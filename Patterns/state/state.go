package main

import "fmt"

// Интерфейс состояния
type State interface {
	handle()
}

// Состояние "Спокойное"
type CalmState struct{}

func (s *CalmState) handle() {
	fmt.Println("Персонаж находится в спокойном состоянии")
}

// Состояние "Боевое"
type BattleState struct{}

func (s *BattleState) handle() {
	fmt.Println("Персонаж находится в боевом состоянии")
}

// Контекст персонажа
type Character struct {
	state State
}

func (c *Character) setState(state State) {
	c.state = state
}

func (c *Character) handle() {
	c.state.handle()
}

func main() {
	character := &Character{}
	calmState := &CalmState{}
	battleState := &BattleState{}

	character.setState(calmState)
	character.handle() // Персонаж находится в спокойном состоянии

	character.setState(battleState)
	character.handle() // Персонаж находится в боевом состоянии
}
