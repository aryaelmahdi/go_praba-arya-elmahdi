package model

import "fmt"

type Car struct {
	carType string
	fuelin  float64
}

func (c *Car) SetCarType(carT string) {
	c.carType = carT
}

func (c *Car) GetCarType() string {
	return c.carType
}

func (c *Car) SetFuelin(fuel float64) {
	c.fuelin = fuel
}

func (c *Car) GetFuelin() float64 {
	return c.fuelin
}

func (c *Car) EstimatedDistance() string {
	distance := fmt.Sprintf("%.0f", c.fuelin/1.5)
	res := "Car Type : " + c.carType + ", " + "Distance : " + distance
	return res
}
