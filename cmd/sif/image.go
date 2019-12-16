package sif

import (
	"github.com/asuahsahua/advent2019/cmd/common"
)

type SpaceImage struct{
	Height int
	Width int
	Layers []*Layer
}

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