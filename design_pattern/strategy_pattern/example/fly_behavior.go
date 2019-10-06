package main

import "fmt"

type FlyBehavior interface{
	Fly() //how you fly
}

type FlyOpenWings struct{

}

func NewFlyOpenWings() FlyBehavior{
	return new(FlyOpenWings)
}

func (*FlyOpenWings) Fly() {
	fmt.Println("I'm spreading wings to fly")
}

type FlyNoWay struct{

}

func NewFlyNoWay() FlyBehavior{
	return new(FlyNoWay)
}

func (*FlyNoWay) Fly() {
	fmt.Println("I'm can't fly")
}
