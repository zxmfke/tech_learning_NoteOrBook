package main

import "fmt"

type Ducky interface {
	Quack()
	Fly()
}

type GreenDuck struct {
}

func (g *GreenDuck) Quack() {
	fmt.Println("I'm quacking, gua gua")
}

func (g *GreenDuck) Fly() {
	fmt.Println("I'm flying")
}