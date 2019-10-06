package main

import "fmt"

type McdonaldDoubleCheeseburger struct {
	Hamburger
}

func NewMcdonaldDoubleCheeseBurger() *McdonaldDoubleCheeseburger {
	return &McdonaldDoubleCheeseburger{
		Hamburger{
			Ingredient: []string{"bread * 2", "cheese * 4", "beef * 2"},
			Name:       "Mcdonald Double Cheeseburger",
		},
	}
}

func (d *McdonaldDoubleCheeseburger) Bake() {
	fmt.Println("bake bread 0.5 min")
}

type McdonaldBigMac struct {
	Hamburger
}

func NewMcdonaldBigMac() *McdonaldBigMac {
	return &McdonaldBigMac{
		Hamburger{
			Ingredient: []string{"bread * 3", "cheese * 2", "beef * 2", "cabbage"},
			Name:       "Mcdonald Big Mac",
		},
	}
}

func (d *McdonaldBigMac) Bake() {
	fmt.Println("bake bread 1.5 min")
}

type McdonaldUnworthy struct {
	Hamburger
}

func NewMcdonaldUnworthy() *McdonaldUnworthy {
	return &McdonaldUnworthy{
		Hamburger{
			Ingredient: []string{"bread * 2", "cheese * 2", "beef * 3", "sausage * 2"},
			Name:       "Mcdonald Unworthy",
		},
	}
}

func (d *McdonaldUnworthy) Bake() {
	fmt.Println("bake bread 2 min")
}

func (d *McdonaldUnworthy) Package() {
	fmt.Println("package with unworthy package")
}
