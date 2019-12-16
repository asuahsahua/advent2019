package intcode

import (
	."github.com/asuahsahua/advent2019/cmd/common"
)

func OptimizeAmpProgram(prog string) []int {
	bestScore := 0
	bestSlice := make([]int, 5)

	IPermutations([]int{0, 1, 2, 3, 4}, func(perm []int) {
		result := AmpProgramRun(prog, perm)
		if result > bestScore {
			bestScore = result
			copy(bestSlice, perm)
		}
	})

	return bestSlice
}

func OptimizeFeedbackProgram(prog string) []int {
	bestScore := 0
	bestSlice := make([]int, 5)

	IPermutations([]int{5, 6, 7, 8, 9}, func(perm []int) {
		result := FeedbackProgramRun(prog, perm)
		if result > bestScore {
			bestScore = result
			copy(bestSlice, perm)
		}
	})

	return bestSlice
}