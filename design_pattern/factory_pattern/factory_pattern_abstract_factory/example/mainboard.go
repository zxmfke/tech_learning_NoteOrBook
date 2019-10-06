package main

import "fmt"

type Mainboard interface {
	InstallCPU()
}

type GAMainboard struct {
	CpuHoles int
}

func NewGAMainboard() *GAMainboard {
	return &GAMainboard{
		CpuHoles: 888,
	}
}

func (g *GAMainboard) InstallCPU() {
	fmt.Printf("this is GA mainboard has %d cpu holes \n", g.CpuHoles)
}

type MSTMainboard struct {
	CpuHoles int
}

func NewMSTMainboard() *MSTMainboard {
	return &MSTMainboard{
		CpuHoles: 1155,
	}
}

func (m *MSTMainboard) InstallCPU() {
	fmt.Printf("this is MST mainboard has %d cpu holes \n", m.CpuHoles)
}
