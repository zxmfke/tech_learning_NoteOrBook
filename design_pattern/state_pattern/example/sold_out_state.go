package main

import "fmt"

type SoldOutState struct {
	Machine *SugarMachine
}

func NewSoldOutState(machine *SugarMachine) SoldOutState {
	return SoldOutState{
		Machine: machine,
	}
}

func (s *SoldOutState) insertOneYuan() {
	fmt.Println("已经售完，不能再投钱了")
}

func (s *SoldOutState) rejectOneYuan() {
	fmt.Println("不能退钱，你还没投钱")
}

func (s *SoldOutState) turnCrank() {
	fmt.Println("糖果已经售空啦")
}

func (s *SoldOutState) dispense() {
	fmt.Println("没糖果出了")
}
