package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func main() {
	input := `hello this is an input`
	charcount := len(input)
	Part1("%s", charcount)
	firstchar := input[0]
	Part2("%d", firstchar)
}