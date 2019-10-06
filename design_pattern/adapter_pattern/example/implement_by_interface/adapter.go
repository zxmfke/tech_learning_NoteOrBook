package main

type Adapter struct {
	Turkey Turkey
}

func NewAdapter() *Adapter{
	return &Adapter{
		Turkey:&WildTurkey{},
	}
}

func (a *Adapter) Quack() {
	a.Turkey.Gobble()
}

func (a *Adapter) Fly() {
	a.Turkey.Fly()
}


