package mediator

type StationManager struct {
	TrainQueue []Train
	isFreeSlot bool
}

func NewStationManager() *StationManager {
	return &StationManager{
		isFreeSlot: true,
	}
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.isFreeSlot {
		s.isFreeSlot = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDepart() {
	if !s.isFreeSlot {
		s.isFreeSlot = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrainInQueue := s.TrainQueue[0]
		firstTrainInQueue.PermitArrival()
	}
}
