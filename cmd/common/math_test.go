package common

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestDecimalDigits(t *testing.T) {
	Equal(t, []int{1, 2, 3, 4, 5, 6}, DecimalDigits(123456))
}

func TestDecimalDigitsReverse(t *testing.T) {
	Equal(t, []int{6, 5, 4, 3, 2, 1}, DecimalDigitsReverse(123456))
}