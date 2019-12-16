package main

import (
	"github.com/asuahsahua/advent2019/cmd/universe"
	"testing"
	. "github.com/stretchr/testify/assert"
)
func TestConfirmedAnswer(t *testing.T) {
	uni := universe.ParseUniverse(myMapData)
	Equal(t, 130681, uni.SumOrbits())
	Equal(t, 313, uni.CountTransfersBetween("YOU", "SAN"))
}