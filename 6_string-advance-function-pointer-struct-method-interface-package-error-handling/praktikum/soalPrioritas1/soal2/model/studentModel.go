package model

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s *Student) SetName(name []string) {
	s.name = name
}

func (s *Student) GetName() []string {
	return s.name
}

func (s *Student) SetScore(score []int) {
	s.score = score
}

func (s *Student) GetScore() []int {
	return s.score
}

func (s *Student) ShowAll() {
	var max, min, ave int
	min = 100
	var lowest, highest string
	for i := 0; i < len(s.score); i++ {
		if s.score[i] > max {
			max = s.score[i]
			highest = s.name[i]
		}
		if s.score[i] < min {
			min = s.score[i]
			lowest = s.name[i]
		}
		ave = ave + s.score[i]
	}

	ave = ave / len(s.score)

	fmt.Println("Average Score :", ave)
	fmt.Println("Min Score of Students :", lowest, "(", min, ")")
	fmt.Println("Max Score of Students :", highest, "(", max, ")")
}

// func (s *Student) AddName(name string) {
// 	s.name = append(s.name, name)
// }

// func (s *Student) AddScore(score int) {
// 	s.score = append(s.score, score)
// }
