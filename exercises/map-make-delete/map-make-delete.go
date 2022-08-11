package main

import "fmt"

func main() {
	carMap := make(map[string]string)

	carMap["Audi"] = "A5"
	carMap["ToyoTa"] = "Yarris"

	value, check := carMap["VW"]
	delete(carMap, "ToyoTa")
	fmt.Printf("%t: %t\n", value == "", check)
	fmt.Println(carMap)

}
