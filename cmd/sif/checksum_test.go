package sif

import (
	"github.com/asuahsahua/advent2019/cmd/common"
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	// Layer1   
	// 021
	// 200
	// Layer 2
	// 110
	// 221
	data := common.DecimalDigitsStr(`021200110221`)
	image := NewSpaceImage(3, 2, data)
	Equal(t, []int{0,2,1,2,0,0}, image.Layers[0].Bytes)
	Equal(t, []int{1,1,0,2,2,1}, image.Layers[1].Bytes)
	// fewest zeroes is layer 2, 3 * 2 = 6
	Equal(t, 6, image.Checksum())
}