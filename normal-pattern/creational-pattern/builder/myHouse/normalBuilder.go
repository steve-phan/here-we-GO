package builder

type normalBuilder struct {
	windowType string
	doorType   string
	numOfFloor int
}

// A function to create an initial normalBuilder
// func newNormalBuilder() *normalBuilder {
// 	return &normalBuilder{}
// }

// A function to customize the type of window
func (b *normalBuilder) setWindowType(wt string) {
	b.windowType = wt + "NORMAL"
}

// A function to customize the type of door
func (b *normalBuilder) setDoorType(dt string) {
	b.doorType = dt
}

// A function to customize the number of floor
func (b *normalBuilder) setNumOfFloor(n int) {
	b.numOfFloor = n
}

func (b *normalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		numOfFloor: b.numOfFloor,
	}
}
