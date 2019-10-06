package main

type Coffee interface {
	Cost() float64
	GetDescription() string
}

type Espresso struct {
	Description string
} // 意式浓缩

func NewEspresso() *Espresso {
	return &Espresso{
		Description: "Espresso",
	}
}

func (e *Espresso) Cost() float64 {
	return 0.4
}

func (e *Espresso) GetDescription() string {
	return "this is Espresso"
}

type DarkRoast struct {
	Description string
} //深度烘焙咖啡

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{
		Description: "DarkRoast",
	}
}

func (e *DarkRoast) Cost() float64 {
	return 0.5
}

func (e *DarkRoast) GetDescription() string {
	return "this is DarkRoast"
}

type Decaf struct {
	Description string
} //低咖啡因咖啡

func NewDecaf() *Decaf {
	return &Decaf{
		Description: "Decaf",
	}
}

func (e *Decaf) Cost() float64 {
	return 0.3
}

func (e *Decaf) GetDescription() string {
	return "this is Decaf"
}