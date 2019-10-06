package main

import "fmt"

type Tea struct {
	CaffeineBeverage
}

func (t *Tea) Brew() {
	fmt.Println("沸水浸泡茶叶")
}

func (t *Tea) AddCondiments() {
	fmt.Println("加柠檬")
}