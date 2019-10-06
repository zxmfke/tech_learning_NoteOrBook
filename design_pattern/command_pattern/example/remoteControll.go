package main

type RemoteController struct{
	Command
}

func  (r *RemoteController) setCommand(command Command) {
	r.Command = command
}

func (r *RemoteController) buttonPress() {
	r.Command.execute()
}
