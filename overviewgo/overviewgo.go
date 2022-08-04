package overviewgo

import (
	"fmt"
	"log"
	// "sort"
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
	// var mySlice []string

	myArray := [10]string{"One", "Two", "Three", "Eins", "Zwei", "Drei", "Vier"}

	// mySlice = append(mySlice, "One", "Two", "Three", "Eins", "Zwei", "Drei", "Vier")

	// Create a Slice from an Array
	myNewSlice := myArray[4:6]

	// sort.Slice(mySlice, func(i, j int) bool {
	// 	return i - j
	// })
	// // mySlice[2] = "Two"
	// sort.Slice(mySlice, func(i, j int) bool {
	// 	return i >= j
	// })
	// sort.Strings(mySlice)

	// log.Println((mySlice))
	// log.Println((mySlice[len(mySlice)-1]))
	myArray[9] = "Nine"
	log.Println("My new Slice is ", myNewSlice, myArray, myNewSlice[0])

}

// Interfaces are implemented implicitly

type Animal interface {
	// Name string;
	SayHi() string
	SayBye() string
}

type Dog struct {
	Name  string
	Breed string
	Legs  int
}

func ShowAnimalsInfo(a Animal) {
	fmt.Println("Animal say Hi ", a.SayHi(), "Animal say Bye ", a.SayBye())
}

func ShowInfo() {
	myDog := Dog{
		Name:  "Cau Vang",
		Breed: "Phu Quoc",
		Legs:  4,
	}
	ShowAnimalsInfo(myDog)

	fmt.Println("my Dog issssssss ", myDog.Breed)

}

// These 2 methods means type Dog implements the interface Animal,
// but we don't need to explicitly declare that it does so.
func (t Dog) SayHi() string {
	return "Woh woh"
}

func (t Dog) SayBye() string {
	return "Grrrrrrrrrrrrrr"
}
