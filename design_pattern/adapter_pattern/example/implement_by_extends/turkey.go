package main

import "fmt"

type Turkey interface{
	Gobble()
	TurkeyFly()
}

type WildTurkey struct{

}

func (w *WildTurkey) Gobble() {
	fmt.Println("I'm gobbling")
}

func (w *WildTurkey) TurkeyFly() {
	fmt.Println("I'm flying, but I can not fly far away")
}