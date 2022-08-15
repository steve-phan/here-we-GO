package mediator

type StationManager struct {
	trainQueue []Train
	isFreeSlot bool
}

func NewStationManager() *StationManager {
	return &StationManager{
		isFreeSlot: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isFreeSlot {
		s.isFreeSlot = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDepart() {
	if !s.isFreeSlot {
		s.isFreeSlot = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		firstTrainInQueue.PermitArrival()
	}
}
