package main

import (
	"regexp"
	"strconv"
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func main() {
	// --- Day 4: Secure Container ---

	// You arrive at the Venus fuel depot only to discover it's protected by a
	// password. The Elves had written the password on a sticky note, but
	// someone threw it out.

	// However, they do remember a few key facts about the password:

	//     It is a six-digit number.
	//     The value is within the range given in your puzzle input.
	//     Two adjacent digits are the same (like 22 in 122345).
	//     Going from left to right, the digits never decrease; they only ever
	//         increase or stay the same (like 111123 or 135679).

	// Other than the range rule, the following are true:

	//     111111 meets these criteria (double 11, never decreases).
	//     223450 does not meet these criteria (decreasing pair of digits 50).
	//     123789 does not meet these criteria (no double).

	// How many different passwords within the range given in your puzzle input
	// meet these criteria?

	// Your puzzle input is 240920-789857.
	input := parseInput1(`240920-789857`)

	Part1("%d", GetPasswordCount(input, false))

	// --- Part Two ---

	// An Elf just remembered one more important detail: the two adjacent
	// matching digits are not part of a larger group of matching digits.

	// Given this additional criterion, but still ignoring the range rule, the
	// following are now true:
	//     112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
	//     123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
	//     111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).

	// How many different passwords within the range given in your puzzle input
	// meet all of the criteria?
	Part2("%d", GetPasswordCount(input, true))
}

func parseInput1(input string) IntRange {
	regexp := regexp.MustCompile(`^(\d+)-(\d+)$`)
	match := regexp.FindStringSubmatch(input)
	PanicIf(match == nil, "Could not find match")

	first, err1 := strconv.Atoi(match[1])
	PanicIfErr(err1)

	second, err2 := strconv.Atoi(match[2])
	PanicIfErr(err2)

	return IntRange{first, second}
}

type IntRange struct{
	start int
	end int
}

func MeetsPasswordCriteria(password int, valueRange IntRange, checkLargerGroup bool) bool {
	// It is a six-digit number.
	if password < 100000 || password > 999999 {
		return false
	}

	// The value is within the range given in your puzzle input.
	if password < valueRange.start || password > valueRange.end {
		return false
	}

	digits := DecimalDigits(password)

	for i := 0; i < len(digits) - 1; i++ {
		first, second := digits[i], digits[i+1]

		// Going from left to right, the digits never decrease; they only ever
		// increase or stay the same (like 111123 or 135679).
		if first > second {
			return false
		}
	}

	adjacents := FindAdjacentGroups(digits)

	if false == checkLargerGroup {
		return len(adjacents) > 0
	}

	for _, group := range(adjacents) {
		if group.count == 2 {
			return true
		}
	}

	// If the previous for{} loop didn't return true, this is an invalid password.
	return false
}

func GetPasswordCount(valueRange IntRange, checkLargerGroup bool) int {
	passwordCount := 0
	for i := valueRange.start ; i <= valueRange.end; i++ {
		if MeetsPasswordCriteria(i, valueRange, checkLargerGroup) {
			passwordCount++
		}
	}
	return passwordCount
}

type AdjacentDigits struct{
	digit int // The digit that has adjacency
	count int // How many digits are adjacent in this match
}

func FindAdjacentGroups(digits []int) []AdjacentDigits {
	adjacent := make([]AdjacentDigits, 0)

	current := -1 // A digit can't be -1, so this is our 'null' value
	count := 1
	for i := 0; i < len(digits); i++ {
		if digits[i] == current {
			// If we have a match, increment the count and move on
			count++
		} else if count > 1 {
			// If we didn't have a match and the count is more than 1, then we have adjacent values!
			// Record those adjacent values and 
			adjacent = append(adjacent, AdjacentDigits{current, count})
			count = 1
		}

		current = digits[i]
	}
	
	// Finally, clean up the tail end of lingering possibilities
	if count > 1 {
		adjacent = append(adjacent, AdjacentDigits{current, count})
	}

	return adjacent
}