package controlcommand

type OnCommand struct {
	// Device is Television or something else
	Device IDevice
}

func (o *OnCommand) Excute() {
	o.Device.on()
}
