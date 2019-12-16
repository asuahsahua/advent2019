package sif

type Layer struct{
	Height int
	Width int
	Bytes []int
}

func NewLayer(height, width int, bytes []int) *Layer {
	cpyBytes := make([]int, len(bytes))
	copy(cpyBytes, bytes)

	return &Layer{
		Height: height,
		Width: width,
		Bytes: cpyBytes,
	}
}

// Count the number of times the given value shows up in the layer
func (l *Layer) Count(value int) int {
	count := 0
	for _, b := range l.Bytes {
		if b == value {
			count++
		}
	}

	return count
}