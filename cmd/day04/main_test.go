package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestMeetsPasswordCriteria(t *testing.T) {
    // 111111 meets these criteria (double 11, never decreases).
	Equal(t, true, MeetsPasswordCriteria(111111, IntRange{0, 999999}, false))
    // 223450 does not meet these criteria (decreasing pair of digits 50).
	Equal(t, false, MeetsPasswordCriteria(223450, IntRange{0, 999999}, false))
    // 123789 does not meet these criteria (no double).
	Equal(t, false, MeetsPasswordCriteria(123789, IntRange{0, 999999}, false))
    // Requires six digits
	Equal(t, false, MeetsPasswordCriteria(12345, IntRange{0, 999999}, false))
	Equal(t, false, MeetsPasswordCriteria(1234567, IntRange{0, 999999}, false))
	// In the given range
	Equal(t, false, MeetsPasswordCriteria(123455, IntRange{123456, 654321}, false))
}

func TestMeetsAdditionalPasswordCriteria(t *testing.T) {
	// 112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
	Equal(t, true, MeetsPasswordCriteria(112233, IntRange{0, 999999}, true))
	// 123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
	Equal(t, false, MeetsPasswordCriteria(123444, IntRange{0, 999999}, true))
	// 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
	Equal(t, true, MeetsPasswordCriteria(111122, IntRange{0, 999999}, true))
}