package main

import "fmt"

type Command interface {
	execute()
}

type Light struct {
}

func (l *Light) on() {
	fmt.Println("light is on")
}

func (l *Light) off() {
	fmt.Println("light is off")
}

type LightOnCommand struct {
	Light
}

func newLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		Light: *light,
	}
}

func (l *LightOnCommand) execute() {
	l.on()
}

type LightOffCommand struct {
	Light
}

func newLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		Light: *light,
	}
}

func (l *LightOffCommand) execute() {
	l.off()
}
