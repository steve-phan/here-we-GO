package controlcommand

type OffCommand struct {
	Device IDevice
}

func (o *OffCommand) Excute() {
	o.Device.off()
}
