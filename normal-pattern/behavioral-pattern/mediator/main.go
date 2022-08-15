package main

import "bookable24.de/main/normal-pattern/behavioral-pattern/mediator/mediator"

func main() {
	stationManager := mediator.NewStationManager()

	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}

	freightTrain := &mediator.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()

	passengerTrain.Depart()

}
