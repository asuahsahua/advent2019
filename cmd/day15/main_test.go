package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestConfirmedAnswer(t *testing.T) {
	Equal(t, 282, NewMapper().FindOxygen().Steps)
	Equal(t, 286, NewMapper().UntilAirFilled())
}