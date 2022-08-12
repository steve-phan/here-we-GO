package main

import (
	"fmt"

	"bookable24.de/main/normal-pattern/creational-pattern/abstract-factory/abstractfactory"
)

func main() {

	appleFactory := abstractfactory.NewApple()

	newPhone := appleFactory.MakePhone()
	newTable := appleFactory.MakeTable()

	fmt.Printf("The price of the new Phone %v\n", newPhone.GetPrice())
	fmt.Printf("The new Phone is for Market: %v\n", newPhone.GetMarket())

	fmt.Printf("\nThe price of the new Table  %v", newTable.GetPrice())
	fmt.Printf("\nIs the new Table support 5G?  %v", newTable.IsSupport5G())

}
