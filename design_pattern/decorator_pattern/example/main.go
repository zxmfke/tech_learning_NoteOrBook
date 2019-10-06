package main

import "fmt"

func main() {
	Espresso := NewEspresso()
	fmt.Println("coffee : ", Espresso.GetDescription())
	fmt.Printf("cost : %.2f \n", Espresso.Cost())
	fmt.Println("---------------")

	EspressoWithMilk := NewMilk(NewEspresso())
	fmt.Println("coffee : ", EspressoWithMilk.GetDescription())
	fmt.Printf("cost : %.2f \n", EspressoWithMilk.Cost())
	fmt.Println("---------------")

	DecafWithMilkWithMocha := NewMocha(NewMilk(NewDecaf()))
	fmt.Println("coffee : ", DecafWithMilkWithMocha.GetDescription())
	fmt.Printf("cost : %.2f \n", DecafWithMilkWithMocha.Cost())
	fmt.Println("---------------")

	DarkRoastWithWhip := NewWhip(NewDarkRoast())
	fmt.Println("coffee : ", DarkRoastWithWhip.GetDescription())
	fmt.Printf("cost : %.2f \n", DarkRoastWithWhip.Cost())
	fmt.Println("---------------")

}