package main

type ComputerFactory interface {
	CreateCPU() CPU
	CreateMainboard() Mainboard
}

type FactoryA struct {
}

func (f *FactoryA) CreateCPU() CPU {
	return NewAMDCPU()
}

func (f *FactoryA) CreateMainboard() Mainboard {
	return NewMSTMainboard()
}

type FactoryB struct {
}

func (f *FactoryB) CreateCPU() CPU {
	return NewInterCPU()
}

func (f *FactoryB) CreateMainboard() Mainboard {
	return NewGAMainboard()
}
