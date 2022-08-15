package mediator

import "fmt"

type FreightTrain struct {
	Mediator Mediator
}

func (p *FreightTrain) Arrive() {
	if !p.Mediator.canArrive(p) {
		fmt.Println("There is no slot, waiting ...")
		return
	}
	fmt.Println("The Pasenger Train Arrived")
}

func (p *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	p.Mediator.notifyAboutDepart()
}

func (p *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	p.Arrive()
}
