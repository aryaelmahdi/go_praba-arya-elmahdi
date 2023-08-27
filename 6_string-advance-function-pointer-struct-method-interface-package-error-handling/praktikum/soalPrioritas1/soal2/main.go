package main

import (
	"praktikum/praktikum/soalPrioritas1/soal2/model"
)

func main() {
	student := new(model.Student)

	// student.AddName("Rizky")
	// names := make([]string, 0, 5)
	names := []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"}
	student.SetName(names)

	scores := []int{80, 75, 70, 75, 60}
	student.SetScore(scores)

	student.ShowAll()
}
