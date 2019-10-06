package main

import "fmt"

type Ducky struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior

	Display func()

	Name string
}

func (this *Ducky) SetName(name string) {
	this.Name = name
}

func (this *Ducky) SetFlyBehavior(behavior FlyBehavior) {
	this.FlyBehavior = behavior
}

func (this *Ducky) SetQuackBehavior(behavior QuackBehavior) {
	this.QuackBehavior = behavior
}

func (this *Ducky) Fly() {
	this.FlyBehavior.Fly()
}

func (this *Ducky) Quack() {
	this.QuackBehavior.Quack()
}

func (this *Ducky) ThisMyName() {
	fmt.Printf("This is my name: %s\n", this.Name)
}

type WoodDucky struct {
	Ducky
}

func (*WoodDucky) Display() {
	fmt.Println("I'm made by wood")
}

func NewWoodDucky() *WoodDucky {
	var ducky = new(WoodDucky)
	ducky.QuackBehavior = NewQuackNoWay()
	ducky.FlyBehavior = NewFlyNoWay()
	ducky.SetName("WoodDucky")
	return ducky
}

type GreenHeadDucky struct {
	Ducky
}

func NewGreenHeadDucky() *GreenHeadDucky {
	var ducky = new(GreenHeadDucky)
	ducky.QuackBehavior = NewQuackJiJi()
	ducky.FlyBehavior = NewFlyOpenWings()
	ducky.SetName("GreenHeadDucky")
	return ducky
}

func (*GreenHeadDucky) Display() {
	fmt.Println("I'm Green duck")
}

type CantFlyDucky struct {
	Ducky
}

func NewCantFlyDucky() *CantFlyDucky {
	var ducky = new(CantFlyDucky)
	ducky.QuackBehavior = NewQuackJiJi()
	ducky.FlyBehavior = NewFlyNoWay()
	ducky.SetName("CantFlyDucky")
	return ducky
}

func (*CantFlyDucky) Display() {
	fmt.Println("I'm duck, but I can not fly")
}
