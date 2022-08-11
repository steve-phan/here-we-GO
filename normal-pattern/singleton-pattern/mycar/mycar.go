package mycar

import (
	"fmt"
	"sync"
)

type MyCar interface {
	AddDrive()
}

type car struct {
	year int
}

var (
	instance *car
	once     sync.Once
)

func GetInstance() MyCar {
	// Do callback only 1 times
	once.Do(func() {
		instance = &car{year: 2022}

	})
	// if instance == nil {
	// }
	return instance
}

func (c *car) AddDrive() {
	c.year++
	fmt.Printf("The current year is %v\n", c.year)
}
