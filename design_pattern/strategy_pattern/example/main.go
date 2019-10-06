package main

import "fmt"

func main()  {


	greenducky := NewGreenHeadDucky()
	greenducky.ThisMyName()
	greenducky.Fly()
	greenducky.Quack()
	greenducky.Display()

	fmt.Println("<----------------------------->")

	mutouDucky := NewWoodDucky()
	mutouDucky.ThisMyName()
	mutouDucky.Fly()
	mutouDucky.Quack()
	mutouDucky.Display()

	fmt.Println("--- I want moutou bian smart ---")

	mutouDucky.SetName("smartDucky")
	mutouDucky.ThisMyName()

	fmt.Println("<----------------------------->")

	cantFlyDucky := NewCantFlyDucky()
	cantFlyDucky.ThisMyName()
	cantFlyDucky.Fly()
	cantFlyDucky.Quack()
	cantFlyDucky.Display()

	fmt.Println("--- I want this ducky fly ---")

	cantFlyDucky.SetFlyBehavior(NewFlyOpenWings())
	cantFlyDucky.Fly()
}