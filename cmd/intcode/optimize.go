package intcode

import (
	."github.com/asuahsahua/advent2019/cmd/common"
)

func OptimizeAmpProgram(prog string) []int64 {
	bestScore := int64(0)
	bestSlice := make([]int64, 5)

	I64Permutations([]int64{0, 1, 2, 3, 4}, func(perm []int64) {
		result := AmpProgramRun(prog, perm)
		if result > bestScore {
			bestScore = result
			copy(bestSlice, perm)
		}
	})

	return bestSlice
}

func OptimizeFeedbackProgram(prog string) []int64 {
	bestScore := int64(0)
	bestSlice := make([]int64, 5)

	I64Permutations([]int64{5, 6, 7, 8, 9}, func(perm []int64) {
		result := FeedbackProgramRun(prog, perm)
		if result > bestScore {
			bestScore = result
			copy(bestSlice, perm)
		}
	})

	return bestSlice
}