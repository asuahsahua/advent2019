package sif

// To make sure the image wasn't corrupted during transmission, the Elves
// would like you to find the layer that contains the fewest 0 digits. On
// that layer, what is the number of 1 digits multiplied by the number of 2
// digits?
func (si *SpaceImage) Checksum() int {
	zeroes := 1 << 32 // maxint substitute
	var current *Layer = nil

	for _, layer := range si.Layers {
		layerZeros := layer.Count(0)
		if layerZeros < zeroes {
			zeroes = layerZeros
			current = layer
		}
	}

	return current.Count(1) * current.Count(2)
}