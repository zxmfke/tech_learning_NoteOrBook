package main

type State interface {
	insertOneYuan()

	rejectOneYuan()

	turnCrank()

	dispense()
}
