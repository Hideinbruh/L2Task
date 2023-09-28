package main

import "fmt"

type Command interface {
	Execute()
}

type LightOnCommand struct {
	Light *Light
}

func (c *LightOnCommand) Execute() {
	c.Light.On()
}

type LightOffCommand struct {
	Light *Light
}

func (c *LightOffCommand) Execute() {
	c.Light.Off()
}

type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("Свет выключен")
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) ExecuteCommands() {
	for _, c := range i.commands {
		c.Execute()
	}
}

func main() {
	light := &Light{}

	lightOnCommand := &LightOnCommand{Light: light}
	lightOffCommand := &LightOffCommand{Light: light}

	invoker := &Invoker{}

	invoker.AddCommand(lightOnCommand)
	invoker.AddCommand(lightOffCommand)

	invoker.ExecuteCommands()
}
