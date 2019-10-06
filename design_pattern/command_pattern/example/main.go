package main

import "fmt"

func main() {
	remoteContorller := new(RemoteController)

	fmt.Println("execute light on")
	light := new(Light)
	lightOnCommand := newLightOnCommand(light)

	remoteContorller.setCommand(lightOnCommand)
	remoteContorller.buttonPress()

	fmt.Println("")
	fmt.Println("execute light off")
	lightOffCommand := newLightOffCommand(light)

	remoteContorller.setCommand(lightOffCommand)
	remoteContorller.buttonPress()
}
