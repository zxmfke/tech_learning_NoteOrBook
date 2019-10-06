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

	fmt.Print("Preparing ")
	for _, v := range h.Ingredient {
		fmt.Printf("%s,", v)
	}
	fmt.Println("prepare finish")
}

func (h *Hamburger) Package() {
	fmt.Println("package hamburger with official package")
}
