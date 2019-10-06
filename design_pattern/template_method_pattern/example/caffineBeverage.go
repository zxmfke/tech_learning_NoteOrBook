package main

import "fmt"

type Beverage interface {
	prepareRecipe()

	BoilWater()

	Brew()

	PourIntoCup()

	AddCondiments()
}

type CaffeineBeverage struct {
	b Beverage
}

func NewBeverage(b Beverage) *CaffeineBeverage {
	return &CaffeineBeverage{
		b: b,
	}
}

func (c *CaffeineBeverage) prepareRecipe() {
	c.b.BoilWater()
	c.b.Brew()
	c.b.PourIntoCup()
	c.b.AddCondiments()
}

func (c *CaffeineBeverage) BoilWater() {
	fmt.Println("烧开水")
}

func (c *CaffeineBeverage) PourIntoCup() {
	fmt.Println("将饮料倒进杯子")
}
