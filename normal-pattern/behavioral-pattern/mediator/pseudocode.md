1. Device: (Train) interface

   - Arrive()
   - Depart()
   - PermitArrival()

2. Mediator: interface

   - CanArrive(Train) bool
   - NotifyAboutDepart()

3. Concrete: Mediator (Mediator) struct

   - PassengerTrain
   - FreghtTrain

   _*implement*_ 3 methods: _Arrive_, _Depart_, _PermitArrival_

4. StationManager struct

   - isFreeSlot bool
   - trainQueue []Train

   _*implement*_ 2 methods: _CanArrive(Train) bool_, _NotifyAboutDepart()_
