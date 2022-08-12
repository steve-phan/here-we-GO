package builder

// A director determines which builder is using
type Director struct {
	builder IBuilder
}

// A function to create a new director
func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse(h House) House {
	d.builder.setWindowType(h.windowType)
	d.builder.setDoorType(h.doorType)
	d.builder.setNumOfFloor(h.numOfFloor)
	return d.builder.getHouse()
}
