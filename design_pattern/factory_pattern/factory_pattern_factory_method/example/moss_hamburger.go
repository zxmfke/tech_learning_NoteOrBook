package main

import "fmt"

type MossDoubleCheeseburger struct {
	Hamburger
}

func NewMossDoubleCheeseBurger() *MossDoubleCheeseburger {
	return &MossDoubleCheeseburger{
		Hamburger{
			Ingredient: []string{"bread * 3", "cheese * 4", "beef", "onion"},
			Name:       "Moss Double Cheeseburger",
		},
	}
}

func (d *MossDoubleCheeseburger) Bake() {
	fmt.Println("bake bread 1 min")
}

type MossBigMac struct {
	Hamburger
}

func NewMossBigMac() *MossBigMac {
	return &MossBigMac{
		Hamburger{
			Ingredient: []string{"bread * 4", "cheese * 3", "beef * 3", "onion"},
			Name:       "Moss Big Mac",
		},
	}
}

func (d *MossBigMac) Bake() {
	fmt.Println("bake bread 3 min")
}

type MossUnworthy struct {
	Hamburger
}

func NewMossUnworthy() *MossUnworthy {
	return &MossUnworthy{
		Hamburger{
			Ingredient: []string{"bread * 2", "cheese * 2", "beef * 2"},
			Name:       "Moss Unworthy",
		},
	}
}

func (d *MossUnworthy) Bake() {
	fmt.Println("bake bread 2 min")
}

func (d *MossUnworthy) Package() {
	fmt.Println("package with unworthy package")
}
