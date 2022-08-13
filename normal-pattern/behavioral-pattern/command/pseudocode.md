1. Button ( Needs some functionality => IControlCommand)

   - command (press)

   1. OnButton => onCommand
   1. OffButton => offCommand

2. Device receiver

   - command (excute)

   1. onCommand => excute() => isRunning = true
   2. ofCommand => excute() => isRunning = false

3. Concrete => Television (Needs some functionality => IDevice)
