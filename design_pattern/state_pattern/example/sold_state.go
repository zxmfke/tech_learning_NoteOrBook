package main

import "fmt"

type SoldState struct {
	Machine *SugarMachine
}

func NewSoldState(machine *SugarMachine) SoldState {
	return SoldState{
		Machine: machine,
	}
}

func (s *SoldState) insertOneYuan() {
	fmt.Println("正在出货中")
}

func (s *SoldState) rejectOneYuan() {
	fmt.Println("已经售出，不可退款")
}

func (s *SoldState) turnCrank() {
	fmt.Println("已经出货，不能再转了")
}

func (s *SoldState) dispense() {
	s.Machine.dispense()
	if s.Machine.Count > 0 {
		fmt.Println("出糖果了")
		s.Machine.setState(&s.Machine.NoOneYuanState)
	} else {
		fmt.Println("糖果已售完")
		s.Machine.setState(&s.Machine.SoldOutState)
	}
}
