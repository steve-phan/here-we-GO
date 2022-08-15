package mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) Arrive() {
	if !p.Mediator.canArrive(p) {
		fmt.Println("There is no slot, waiting ...")
		return
	}
	fmt.Println("The Pasenger Train Arrived")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.Mediator.notifyAboutDepart()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.Arrive()
}
