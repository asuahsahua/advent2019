package sif

import (
	"github.com/asuahsahua/advent2019/cmd/common"
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	// The layers are rendered with the first layer in front and the last layer
	// in back. So, if a given position has a transparent pixel in the first and
	// second layers, a black pixel in the third layer, and a white pixel in the
	// fourth layer, the final image would have a black pixel at that position.

	// For example, given an image 2 pixels wide and 2 pixels tall, the image
	// data 0222112222120000
	imageBytes := common.DecimalDigitsStr(`0222112222120000`)
	image := NewSpaceImage(2, 2, imageBytes)

	// Corresponds to the following image layers:
	// Layer 1: 02
	//          22
	// Layer 2: 11
	//          22
	// Layer 3: 22
	//          12
	// Layer 4: 00
	//          00
	// So, the final image looks like this:
	// 01
	// 10
	Equal(t, []int{0,1,1,0}, image.Render())
}