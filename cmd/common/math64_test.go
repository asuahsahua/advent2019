package common

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestDecimalDigits64(t *testing.T) {
	Equal(t, []int64{1, 2, 3, 4, 5, 6}, DecimalDigits64(123456))
}

func TestDecimalDigitsReverse64(t *testing.T) {
	Equal(t, []int64{6, 5, 4, 3, 2, 1}, DecimalDigitsReverse64(123456))
}