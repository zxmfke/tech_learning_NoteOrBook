package main

import "fmt"

func main() {
	hamburgerFactory := new(HamburgerFactory)

	BigMacBurger := hamburgerFactory.CreateHamburger("Big Mac")
	fmt.Println("This is your ", BigMacBurger.GetName())
	fmt.Println("--------------------------------------")

	DCBurger := hamburgerFactory.CreateHamburger("Double Cheeseburger")
	fmt.Println("This is your ", DCBurger.GetName())
	fmt.Println("--------------------------------------")

	UnworthyBurger := hamburgerFactory.CreateHamburger("Unworthy")
	fmt.Println("This is your ", UnworthyBurger.GetName())
	fmt.Println("--------------------------------------")
}
