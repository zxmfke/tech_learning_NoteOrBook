package main

import "fmt"

func main() {
	sugarMachine := NewSugarMachine(2)

	fmt.Println("============first round============")
	sugarMachine.insertOneYuan()
	sugarMachine.turnCrank()
	fmt.Println("============second round============")
	sugarMachine.insertOneYuan()
	sugarMachine.rejectOneYuan()
	fmt.Println("============third round============")
	sugarMachine.insertOneYuan()
	sugarMachine.turnCrank()
	fmt.Println("============fourth round============")
	sugarMachine.insertOneYuan()
	sugarMachine.turnCrank()
}
