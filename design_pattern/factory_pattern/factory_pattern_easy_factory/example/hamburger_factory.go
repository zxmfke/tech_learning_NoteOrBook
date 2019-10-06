package main

import "fmt"

type HamburgerBasic interface {
	Prepare()
	Bake()
	Package()
	GetName() string
}

type Hamburger struct {
	HamburgerBasic

	Ingredient []string

	Name string
}

func (d *Hamburger) GetName() string {
	return d.Name
}

func (h *Hamburger) Prepare() {
	h.Ingredient = append(h.Ingredient, "bread")
	fmt.Println("add ingredient bread")

	h.Ingredient = append(h.Ingredient, "beef")
	fmt.Println("add ingredient beef")

	h.Ingredient = append(h.Ingredient, "cabbage")
	fmt.Println("add ingredient cabbage")
}

func (h *Hamburger) Bake() {
	fmt.Println("bake bread 1 min")
}

func (h *Hamburger) Package() {
	fmt.Println("package hamburger with official package")
}

type HamburgerFactory struct {

}
func (h *HamburgerFactory) CreateHamburger(typ string) HamburgerBasic {

	var hamburger HamburgerBasic

	switch (typ) {
	case "Double Cheeseburger":
		hamburger = NewDoubleCheeseBurger()
		break
	case "Big Mac":
		hamburger = NewBigMac()
		break
	case "Unworthy":
		hamburger = NewUnworthy()
		break
	default:
		return nil
	}

	hamburger.Prepare()
	hamburger.Bake()
	hamburger.Package()
	return hamburger
}
