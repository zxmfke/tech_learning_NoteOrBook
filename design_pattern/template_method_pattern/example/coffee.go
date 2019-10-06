package main

import "fmt"

type Coffee struct {
	CaffeineBeverage
}

func (c *Coffee) Brew() {
	fmt.Println("冲泡咖啡粉")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("加糖和牛奶")
}