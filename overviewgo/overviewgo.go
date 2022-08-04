package overviewgo

import (
	"log"
	"sort"
)

type TCar struct {
	Brand string
	Mode  string
	Year  int
}

func ShowMyMap() {
	myMap := make(map[string]TCar)

	myCar := TCar{
		Brand: "Audi",
		Mode:  "A2",
		Year:  2020,
	}

	myWifeCar :=
		TCar{
			Brand: "VW",
			Mode:  "TT",
			Year:  2018,
		}
	myMap["myCar"] = myCar
	myMap["myWifeCar"] = myWifeCar

	log.Println(myMap)
}

func ShowMySlice() {
	var mySlice []string
	mySlice = append(mySlice, "One", "Two", "Three", "Eins")

	// sort.Slice(mySlice, func(i, j int) bool {
	// 	return i - j
	// })
	// // mySlice[2] = "Two"
	// sort.Slice(mySlice, func(i, j int) bool {
	// 	return i >= j
	// })
	sort.Strings(mySlice)

	log.Println((mySlice))
	log.Println((mySlice[len(mySlice)-1]))

}
