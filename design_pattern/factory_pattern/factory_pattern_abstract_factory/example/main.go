package main

import "fmt"

func main() {
	engineer := new(Engineer)
	factoryA := new(FactoryA)
	engineer.MakeComputer(factoryA)
	fmt.Println("----------------")


	factoryB := new(FactoryB)
	engineer.MakeComputer(factoryB)
	fmt.Println("----------------")

}