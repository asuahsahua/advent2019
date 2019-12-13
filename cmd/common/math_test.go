package common

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestDecimalDigits(t *testing.T) {
	Equal(t, []int{1, 2, 3, 4, 5, 6}, DecimalDigits(123456))
}