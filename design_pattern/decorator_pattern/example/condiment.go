package main

type Condiment struct {
	Coffee
}

func NewCondiment(coffee Coffee) *Condiment {
	return &Condiment{coffee}
}

func (c *Condiment) Cost() float64 {
	return 0.1
}

func (c *Condiment) GetDescription() string {
	return c.Coffee.GetDescription()
}

type Milk struct {
	*Condiment
}

func NewMilk(coffee Coffee) *Milk {
	this := new(Milk)
	this.Condiment = NewCondiment(coffee)
	return this
}

func (m *Milk) GetDescription() string {
	return m.Coffee.GetDescription() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return 0.2 + m.Coffee.Cost()
}

type Mocha struct {
	*Condiment
}

func NewMocha(coffee Coffee) *Mocha {
	this := new(Mocha)
	this.Condiment = NewCondiment(coffee)
	return this
}

func (m *Mocha) GetDescription() string {
	return m.Coffee.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float64 {
	return 0.3 + m.Coffee.Cost()
}

type Whip struct {
	*Condiment
}

func NewWhip(coffee Coffee) *Whip {
	this := new(Whip)
	this.Condiment = NewCondiment(coffee)
	return this
}

func (w *Whip) GetDescription() string {
	return w.Coffee.GetDescription() + ", Whip"
}

func (w *Whip) Cost() float64 {
	return 0.1 + w.Coffee.Cost()
}
