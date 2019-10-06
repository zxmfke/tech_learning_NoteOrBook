package main

import "fmt"

type CPU interface {
	Calculate()
}

type IntelCPU struct {
	Pins int
}

func NewInterCPU() *IntelCPU {
	return &IntelCPU{
		Pins: 888,
	}
}

func (i *IntelCPU) Calculate() {
	fmt.Printf("this is intel cpu has %d pings \n", i.Pins)
}


type AMDCPU struct {
	Pins int
}

func NewAMDCPU() *AMDCPU {
	return &AMDCPU{
		Pins: 1155,
	}
}

func (m *AMDCPU) Calculate() {
	fmt.Printf("this is AMD cpu has %d pings \n", m.Pins)
}
