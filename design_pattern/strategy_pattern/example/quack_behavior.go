package main

import "fmt"

type QuackBehavior interface{
	Quack() //how you fly
}

type QuackGuaGua struct{

}

func NewQuackGuaGua() QuackBehavior{
	return new(QuackGuaGua)
}

func (*QuackGuaGua) Quack() {
	fmt.Println("I'm quacking: Gua Gua")
}

type QuackNoWay struct{

}

func NewQuackNoWay() QuackBehavior{
	return new(QuackNoWay)
}

func (*QuackNoWay) Quack() {
	fmt.Println("I'm can't Quack")
}

type QuackJiJi struct {

}

func NewQuackJiJi() QuackBehavior{
	return new(QuackJiJi)
}

func (*QuackJiJi) Quack() {
	fmt.Println("I'm quacking: Ji Ji")
}