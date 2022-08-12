package main

import (
	"fmt"

	"bookable24.de/main/normal-pattern/creational-pattern/abstract-factory/abstractfactory"
)

func main() {

	appleFactory := abstractfactory.NewApple()

	newPhone := appleFactory.MakePhone()
	fmt.Printf("Hello, World %v", newPhone.GetPrice())

}
