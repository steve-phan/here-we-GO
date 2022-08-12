package builder

type iglooBuilder struct {
	windowType string
	doorType   string
	numOfFloor int
}

// A function create an initial iglooBuilder
// func newIglooBuilder() *iglooBuilder {
// 	return &iglooBuilder{}
// }

func (b *iglooBuilder) setWindowType(wt string) {
	b.windowType = wt
}

func (b *iglooBuilder) setDoorType(dt string) {
	b.doorType = dt
}

func (b *iglooBuilder) setNumOfFloor(n int) {
	b.numOfFloor = n
}

func (b *iglooBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		numOfFloor: b.numOfFloor,
	}
}
