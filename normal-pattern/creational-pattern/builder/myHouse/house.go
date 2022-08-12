package builder

// Basically provide some getters of House package

type House struct {
	windowType string
	doorType   string
	numOfFloor int
}

func NewHouse() *House {
	return &House{}
}

func (h *House) SetInitialHouse(w string, d string, n int) {
	h.windowType = w
	h.doorType = d
	h.numOfFloor = n

}

func (h *House) GetWindowType() string {
	return h.windowType
}

func (h *House) GetDoorType() string {
	return h.doorType
}

func (h *House) GetNumOfFloor() int {
	return h.numOfFloor
}
