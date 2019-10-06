package main

type Engineer struct {
	Cpu       CPU
	Mainboard Mainboard
}

func (e *Engineer) MakeComputer(factory ComputerFactory) {
	e.prepareHardwares(factory)
}

func (e *Engineer) prepareHardwares(factory ComputerFactory) {
	e.Cpu = factory.CreateCPU()
	e.Mainboard = factory.CreateMainboard()

	e.Cpu.Calculate()
	e.Mainboard.InstallCPU()
}
