package main

import "fmt"

type DoubleCheeseburger struct {
	Hamburger
}

func NewDoubleCheeseBurger() *DoubleCheeseburger {
	return &DoubleCheeseburger{
		Hamburger{
			Ingredient: make([]string, 0),
			Name:       "Double Cheeseburger",
		},
	}
}

func (d *DoubleCheeseburger) Prepare() {
	d.Ingredient = append(d.Ingredient, "bread * 2")
	fmt.Println("add ingredient bread * 2")

	d.Ingredient = append(d.Ingredient, "beef")
	fmt.Println("add ingredient beef")

	d.Ingredient = append(d.Ingredient, "cabbage")
	fmt.Println("add ingredient cabbage")

	d.Ingredient = append(d.Ingredient, "cheese * 2")
	fmt.Println("add ingredient cheese * 2")
}

func (d *DoubleCheeseburger) Bake() {
	fmt.Println("bake bread 0.5 min")
}

type BigMac struct {
	Hamburger
}

func NewBigMac() *BigMac {
	return &BigMac{
		Hamburger{
			Ingredient: make([]string, 0),
			Name:       "Big Mac",
		},
	}
}

func (d *BigMac) Prepare() {
	d.Ingredient = append(d.Ingredient, "bread * 3")
	fmt.Println("add ingredient bread * 3")

	d.Ingredient = append(d.Ingredient, "beef * 2")
	fmt.Println("add ingredient beef * 2")

	d.Ingredient = append(d.Ingredient, "cabbage")
	fmt.Println("add ingredient cabbage")

	d.Ingredient = append(d.Ingredient, "cheese * 2")
	fmt.Println("add ingredient cheese * 2")
}

func (d *BigMac) Bake() {
	fmt.Println("bake bread 1.5 min")
}

type Unworthy struct {
	Hamburger
}

func NewUnworthy() *Unworthy {
	return &Unworthy{
		Hamburger{
			Ingredient: make([]string, 0),
			Name:       "Unworthy",
		},
	}
}

func (d *Unworthy) Prepare() {
	d.Ingredient = append(d.Ingredient, "bread * 2")
	fmt.Println("add ingredient bread * 2")

	d.Ingredient = append(d.Ingredient, "beef * 2")
	fmt.Println("add ingredient beef * 2")

	d.Ingredient = append(d.Ingredient, "sausage * 2")
	fmt.Println("add ingredient sausage * 2")

	d.Ingredient = append(d.Ingredient, "cheese")
	fmt.Println("add ingredient cheese")
}

func (d *Unworthy) Bake() {
	fmt.Println("bake bread 2 min")
}

func (d *Unworthy) Package() {
	fmt.Println("package with unworthy package")
}
