package main

type Adapter struct {
	WildTurkey
	Ducky
}

func NewAdapter() *Adapter{
	return &Adapter{}
}

func (a *Adapter) Quack() {
	a.Gobble()
}

func (a *Adapter) Fly() {
	a.TurkeyFly()
}