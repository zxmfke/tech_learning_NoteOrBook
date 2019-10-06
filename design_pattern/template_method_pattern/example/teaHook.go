package main

import "fmt"

type TeaHook struct {
	CaffeineBeverageHook
}

func (t *TeaHook) Brew() {
	fmt.Println("沸水浸泡茶叶")
}

func (t *TeaHook) AddCondiments() {
	fmt.Println("加柠檬")
}

func (t *TeaHook) CustomerWantsCondiments() bool {
	var answer string
	fmt.Print("是否需要添加柠檬(yes/no):")
	fmt.Scanf("%s\n",&answer)
	return answer == "yes"
}
