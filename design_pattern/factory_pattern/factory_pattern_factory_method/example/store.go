package main

import (
	"fmt"
)

type StoreBasic interface {
	OrderHamburger(string)
	CreateHamburger(string) HamburgerBasic
}

type Mcdonald struct {
	StoreBasic
}

func (m *Mcdonald) OrderHamburger(typ string)  {
	hamburger := m.CreateHamburger(typ)
	hamburger.Prepare()
	hamburger.Bake()
	hamburger.Package()
	fmt.Printf("This is your %s\n", hamburger.GetName())
	fmt.Println("--------------------------------------")
}

func (m *Mcdonald) CreateHamburger(typ string) HamburgerBasic {
	switch (typ) {
	case "Double Cheeseburger":
		return NewMcdonaldDoubleCheeseBurger()
	case "Big Mac":
		return NewMcdonaldBigMac()
	case "Unworthy":
		return NewMcdonaldUnworthy()
	default:
		return nil
	}
}

type Moss struct {
	StoreBasic
}

func (m *Moss) OrderHamburger(typ string) HamburgerBasic {
	hamburger := m.CreateHamburger(typ)
	hamburger.Prepare()
	hamburger.Bake()
	hamburger.Package()
	fmt.Printf("This is your %s\n", hamburger.GetName())
	fmt.Println("--------------------------------------")
	return hamburger
}

func (m *Moss) CreateHamburger(typ string) HamburgerBasic {
	switch (typ) {
	case "Double Cheeseburger":
		return NewMossDoubleCheeseBurger()
	case "Big Mac":
		return NewMossBigMac()
	case "Unworthy":
		return NewMossUnworthy()
	default:
		return nil
	}
}
