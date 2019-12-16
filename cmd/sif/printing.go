package sif

import (
	"fmt"
)

func (si *SpaceImage) PrintRender() {
	rendered := si.Render()

	for y := 0; y < si.Height; y++ {
		for x := 0; x < si.Width; x++ {
			switch rendered[y * si.Width + x] {
			case C_TRANSPARENT:
				fmt.Printf("%s", " ")
				continue
			case C_WHITE:
				fmt.Printf("%s", "B")
			case C_BLACK:
				fmt.Printf("%s", " ")
			default:
				panic("unknown color code")
			}
		}
		fmt.Println()
	}
}