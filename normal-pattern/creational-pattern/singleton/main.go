package main

import (
	"time"

	"bookable24.de/main/normal-pattern/singleton-pattern/mycar"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			car := mycar.GetInstance()
			car.AddDrive()
			car.AddDrive()
			car.AddDrive()
			car.AddDrive()

		}()
	}
	time.Sleep(time.Second)
}
