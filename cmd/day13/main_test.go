package main

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestConfirmedAnswer(t *testing.T) {
	cab := NewArcadeCabinet()
	cab.Run()
	Equal(t, 320, cab.CountTiles(BlockTile))

	// Part 2
	Equal(t, 15156, NewArcadeCabinet().BeatGame())
}
