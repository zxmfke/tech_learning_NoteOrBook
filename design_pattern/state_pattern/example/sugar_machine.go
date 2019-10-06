package main

type SugarMachine struct {
	HasOneYuanState HasOneYuanState
	NoOneYuanState  NoOneYuanState
	SoldOutState    SoldOutState
	SoldState       SoldState

	CurrentState State
	Count        int
}

func NewSugarMachine(count int) *SugarMachine {
	machine := &SugarMachine{
		Count:count,
	}
	hasOneYuanState := NewHasOneYuanState(machine)
	noOneYuanState := NewNoOneYuanState(machine)
	soldOutState := NewSoldOutState(machine)
	soldState := NewSoldState(machine)

	machine.setState(&soldOutState)
	if count > 0 {
		machine.setState(&noOneYuanState)
	}

	machine.NoOneYuanState = noOneYuanState
	machine.SoldState = soldState
	machine.SoldOutState = soldOutState
	machine.HasOneYuanState = hasOneYuanState

	return machine
}

func (s *SugarMachine) insertOneYuan() {
	s.CurrentState.insertOneYuan()
}

func (s *SugarMachine) rejectOneYuan() {
	s.CurrentState.rejectOneYuan()
}

func (s *SugarMachine) turnCrank() {
	s.CurrentState.turnCrank()
	s.CurrentState.dispense()
}

func (s *SugarMachine) dispense() {
	if s.Count > 0 {
		s.Count -= 1
	}
}

func (s *SugarMachine) setState(state State) {
	s.CurrentState = state
}
