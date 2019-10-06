package main

import "fmt"

type HasOneYuanState struct {
	State
	Machine *SugarMachine
}

func NewHasOneYuanState(machine *SugarMachine) HasOneYuanState {
	return HasOneYuanState{
		Machine: machine,
	}
}

func (s *HasOneYuanState) insertOneYuan() {
	fmt.Println("没有投钱")
}

func (s *HasOneYuanState) rejectOneYuan() {
	fmt.Println("退钱中")
	s.Machine.setState(&s.Machine.NoOneYuanState)
}

func (s *HasOneYuanState) turnCrank() {
	fmt.Println("出货中")
	s.Machine.setState(&s.Machine.SoldState)
}

func (s *HasOneYuanState) dispense() {
	fmt.Println("服务器错误")
}
