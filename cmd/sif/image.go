package sif

import (
	"github.com/asuahsahua/advent2019/cmd/common"
)

type SpaceImage struct{
	Height int
	Width int
	Layers []*Layer
}

// The digits indicate the color of the corresponding pixel: 0 is black, 1 is
// white, and 2 is transparent.
const (
	C_BLACK = 0
	C_WHITE = 1
	C_TRANSPARENT = 2
)

func NewSpaceImage(height, width int, bytes []int) *SpaceImage {
	layerSize := height * width
	common.PanicIf(len(bytes) % layerSize != 0, "Expected bytes to be a multiple of the layer size")
	layerCount := len(bytes) / layerSize

	layers := make([]*Layer, 0)
	for i := 0; i < layerCount; i++ {
		layers = append(layers, NewLayer(
			height,
			width,
			bytes[i * layerSize:(i+1) * layerSize],
		))
	}

	return &SpaceImage{
		Height: height,
		Width: width,
		Layers: layers,
	}
}