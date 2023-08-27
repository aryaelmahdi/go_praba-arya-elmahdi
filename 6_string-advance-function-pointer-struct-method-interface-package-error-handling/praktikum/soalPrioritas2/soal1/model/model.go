package model

import (
	"fmt"
)

type ASCI struct {
	input  string
	offset int
}

func (a *ASCI) SetInput(input string) {
	a.input = input
}

func (a *ASCI) GetInput() string {
	return a.input
}

func (a *ASCI) SetOffset(input int) {
	a.offset = input
}

func (a *ASCI) GetOffset() int {
	return a.offset
}

func (a *ASCI) ConvertASCII(input string, offset int) string {
	var ascii string
	a.SetInput(input)
	a.SetOffset(offset)
	for _, character := range a.GetInput() {
		modOff := a.GetOffset() % 26
		tmp := int(character) + modOff
		if tmp > 122 {
			tmp -= 26
		}
		ascii += fmt.Sprintf("%c", tmp)
	}
	return ascii
}
