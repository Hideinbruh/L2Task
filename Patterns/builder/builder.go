package main

import "fmt"

type Computer struct {
	CPU     string
	RAM     int
	GPU     string
	HDD     int
	SSD     int
	Monitor string
}

type ComputerBuilder struct {
	comp Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{comp: Computer{}}
}

func (c *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	c.comp.CPU = cpu
	return c
}

func (c *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	c.comp.RAM = ram
	return c
}

func (c *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	c.comp.GPU = gpu
	return c
}

func (c *ComputerBuilder) SetHDD(hdd int) *ComputerBuilder {
	c.comp.HDD = hdd
	return c
}

func (c *ComputerBuilder) SetSSD(ssd int) *ComputerBuilder {
	c.comp.SSD = ssd
	return c
}

func (c *ComputerBuilder) SetMonitor(monitor string) *ComputerBuilder {
	c.comp.Monitor = monitor
	return c
}

func (c *ComputerBuilder) Build() Computer {
	return c.comp
}

func main() {
	builder := NewComputerBuilder()
	computer := builder.SetCPU("Inter core i7").SetRAM(16).SetGPU("Nvidia GeForce RTX 3080").SetHDD(2000).SetSSD(500).SetMonitor("27-inch 4K").Build()
	fmt.Println(computer)
}
