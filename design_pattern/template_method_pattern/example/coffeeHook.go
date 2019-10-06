package main

import "fmt"

type CoffeeHook struct {
	CaffeineBeverageHook
}

func (c *CoffeeHook) Brew() {
	fmt.Println("冲泡咖啡粉")
}

func (c *CoffeeHook) AddCondiments() {
	fmt.Println("加糖和牛奶")
}

func (c *CoffeeHook) CustomerWantsCondiments() bool {
	var answer string
	fmt.Print("是否需要添加糖和奶(yes/no):")
	fmt.Scanf("%s\n",&answer)
	return answer == "yes"
}