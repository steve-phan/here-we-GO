package controlcommand

import "fmt"

// Television is an Object with a isRunning state
type Television struct {
	isRunning bool
}

// Need implements on() and off() commands

func (t *Television) on() {
	t.isRunning = true
	fmt.Println("Television is On")
}

func (t *Television) off() {
	t.isRunning = false
	fmt.Println(("Television is Off"))
}

/*
=====>RESULT HERE<=====
 A television with
    1. An isRunning state
	2. Methods: on() and off()
*/
