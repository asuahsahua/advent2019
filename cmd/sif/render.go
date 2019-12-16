package sif

import (
	"github.com/asuahsahua/advent2019/cmd/common"
)

// Render the space image by processing each layer according to transparency rules:
//
// The layers are rendered with the first layer in front and the last layer in
// back. So, if a given position has a transparent pixel in the first and second
// layers, a black pixel in the third layer, and a white pixel in the fourth
// layer, the final image would have a black pixel at that position.
func (si *SpaceImage) Render() []int {
	imageSize := si.Height * si.Width
	resultingImage := make([]int, imageSize)

	for pixel := 0; pixel < imageSize; pixel++ {
		resultingImage[pixel] = si.RenderPixel(pixel)
	}

	return resultingImage
}

func (si *SpaceImage) RenderPixel(pixel int) int {
	for _, layer := range si.Layers {
		switch layer.Bytes[pixel] {
		case C_TRANSPARENT:
			continue
		case C_WHITE:
			return C_WHITE
		case C_BLACK:
			return C_BLACK
		default:
			common.Panic("Unrecognized pixel: %d", layer.Bytes[pixel])
		}

		return layer.Bytes[pixel]
	}

	return C_TRANSPARENT
}