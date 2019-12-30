package common

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestFillRepeat(t *testing.T) {
	Equal(t, []int{1, 2, 3, 1, 2}, FillRepeat([]int{1, 2, 3}, 5))
}
