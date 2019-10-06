package main

import "fmt"

type Turkey interface{
	Gobble()
	Fly()
}

type WildTurkey struct{

}

func (w *WildTurkey) Gobble() {
	fmt.Println("I'm gobbling")
}

func (w *WildTurkey) Fly() {
	fmt.Println("I'm flying, but I can not fly far away")
}