package main

import "fmt"

func main() {
	fmt.Println("======开始煮咖啡=====")
	beverage := NewBeverage(&Coffee{})
	beverage.prepareRecipe()

	fmt.Println("======开始泡茶=====")

	beverage = NewBeverage(&Tea{})
	beverage.prepareRecipe()

	fmt.Println("======开始煮钩子咖啡=====")
	beverageHook := NewBeverageHook(&CoffeeHook{})
	beverageHook.prepareRecipeHook()

	//fmt.Println("======开始泡钩子茶=====")
	//
	//beverageHook = NewBeverageHook(&TeaHook{})
	//beverageHook.prepareRecipeHook()

}