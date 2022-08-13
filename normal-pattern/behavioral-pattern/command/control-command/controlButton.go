package controlcommand

/*
 This is a button (an object)
 Button needs a Press medthod
*/
type ControlButton struct {
	Command IControlCommand
}

func (c *ControlButton) Press() {
	c.Command.Excute()
}
