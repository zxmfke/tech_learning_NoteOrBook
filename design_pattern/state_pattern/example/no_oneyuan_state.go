package main

import "fmt"

type NoOneYuanState struct {
	Machine *SugarMachine
}

func NewNoOneYuanState(machine *SugarMachine) NoOneYuanState {
	return NoOneYuanState{
		Machine: machine,
	}
}

func (n *NoOneYuanState) insertOneYuan() {
	fmt.Println("投了1元")
	n.Machine.setState(&n.Machine.HasOneYuanState)
}

func (n *NoOneYuanState) rejectOneYuan() {
	fmt.Println("你还没投钱，无法退钱")
}

func (n *NoOneYuanState) turnCrank() {
	fmt.Println("还么有投钱，不能转手柄")
}

func (n *NoOneYuanState) dispense() {
	fmt.Println("先付钱啊")
}
