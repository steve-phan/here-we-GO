package builder

type IBuilder interface {
	setWindowType(w string)
	setDoorType(d string)
	setNumOfFloor(n int)
	getHouse() House
}

type TBuilderType string

const (
	NORMALHOUSE TBuilderType = "NORMALHOUSE"
	IGLOOHOUSE  TBuilderType = "IGLOOHOUSE"
)

// A function to get builder to build a house
func GetBuilder(builderType TBuilderType) IBuilder {
	switch builderType {
	case NORMALHOUSE:
		return &normalBuilder{}
	case IGLOOHOUSE:
		return &iglooBuilder{}
	}
	return &normalBuilder{}
}
