package mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) Arrive() {
	if !p.Mediator.CanArrive(p) {
		fmt.Println("There is no slot, waiting ...")
		return
	}
	fmt.Println("The Pasenger Train Arrived")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.Mediator.NotifyAboutDepart()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.Arrive()
}
