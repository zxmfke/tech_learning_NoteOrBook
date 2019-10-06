package main

import "fmt"

func main() {
	var ducky Ducky

	ducky = &GreenDuck{}
	DuckyShow(ducky)

	ducky = NewAdapter()

	DuckyShow(ducky)
}

func DuckyShow(ducky Ducky) {
	ducky.Quack()
	ducky.Fly()
	fmt.Println("---------Finish show---------")
}