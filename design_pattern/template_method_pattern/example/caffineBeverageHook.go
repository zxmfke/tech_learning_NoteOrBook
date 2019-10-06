package main

import "fmt"

type BeverageHook interface {
	prepareRecipeHook()

	BoilWater()

	Brew()

	PourIntoCup()

	AddCondiments()

	CustomerWantsCondiments() bool
}

type CaffeineBeverageHook struct {
	b BeverageHook
}

func NewBeverageHook(b BeverageHook) *CaffeineBeverageHook {
	return &CaffeineBeverageHook{
		b: b,
	}
}

func (c *CaffeineBeverageHook) prepareRecipeHook() {
	c.b.BoilWater()
	c.b.Brew()
	c.b.PourIntoCup()
	if c.b.CustomerWantsCondiments() {
		c.b.AddCondiments()
	} else {
		fmt.Println("不添加调料")
	}
}

func (c *CaffeineBeverageHook) BoilWater() {
	fmt.Println("烧开水")
}

func (c *CaffeineBeverageHook) PourIntoCup() {
	fmt.Println("将饮料倒进杯子")
}

func (c *CaffeineBeverageHook) CustomerWantsCondiments() bool {
	return true
}