package mediator

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDepart()
}
