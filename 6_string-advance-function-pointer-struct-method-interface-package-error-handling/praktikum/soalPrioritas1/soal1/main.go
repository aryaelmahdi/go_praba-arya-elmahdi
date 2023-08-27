package main

import (
	"fmt"
	"praktikum/praktikum/soalPrioritas1/soal1/model"
)

func main() {
	car := new(model.Car)
	var carType = "SUV"
	var fuelin = 10.5

	car.SetCarType(carType)
	car.SetFuelin(fuelin)

	fmt.Println((car.EstimatedDistance()))
}
