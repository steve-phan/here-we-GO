package main

import (
	"fmt"

	builder "bookable24.de/main/normal-pattern/creational-pattern/builder/myHouse"
)

func main() {
	newHouse := builder.NewHouse()
	newHouse.SetInitialHouse("wooden", "steel", 2)
	normailBuilder := builder.GetBuilder(builder.NORMALHOUSE)
	// iglooBuilder := builder.GetBuilder(builder.IGLOOHOUSE)

	// Do we need a director here?
	director := builder.NewDirector(normailBuilder)
	// var newHouseInput builder.House

	normalHouse := director.BuildHouse(*newHouse)

	fmt.Printf("The window type of normal House is %s\n", normalHouse.GetWindowType())
	fmt.Printf("The number of floor of normal House is %d\n", normalHouse.GetNumOfFloor())
}
