package main

import (
	cc "bookable24.de/main/normal-pattern/behavioral-pattern/command/control-command"
)

func main() {

	tv := &cc.Television{}

	onCommand := &cc.OnCommand{
		Device: tv,
	}

	onButton := &cc.ControlButton{
		Command: onCommand,
	}
	onButton.Press()

	offCommand := &cc.OffCommand{
		Device: tv,
	}

	offButton := &cc.ControlButton{
		Command: offCommand,
	}

	offButton.Press()

}
