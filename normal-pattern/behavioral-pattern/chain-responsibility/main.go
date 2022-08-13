package main

import "bookable24.de/main/normal-pattern/behavioral-pattern/chain-responsibility/chanres"

func main() {
	/*
		while Cashier is a struct, we create a cashier instace
		This cashier also have to methods
	*/

	cashier := &chanres.Cashier{}

	medical := &chanres.Medical{}
	medical.SetNext(cashier)

	doctor := &chanres.Doctor{}
	doctor.SetNext(medical)

	reception := &chanres.Reception{}
	reception.SetNext(doctor)

	// Initial a patient
	patient := &chanres.Patient{Name: "Teddy"}

	// First the patient go to the reception
	reception.Excute(patient)

}
