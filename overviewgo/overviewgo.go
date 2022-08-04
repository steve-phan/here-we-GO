package overviewgo

import "log"

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
	myMap["myCar"] = myCar

	log.Println(myMap["myCar"].Brand)
}
