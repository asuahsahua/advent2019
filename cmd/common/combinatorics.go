package common

// Generate all permutations of the given values. Assumes input is unique.
func IPermutations(A []int, callback func([]int)) {
	// Heap's algorithm https://en.wikipedia.org/wiki/Heap%27s_algorithm
	// Using the non-recursive form, which is going to just look wacky since
	// it manually tracks a stack.

	n := len(A)
	// c is an encoding of the stack state. c[k] encodes the for-loop counter
	// for when generate(k+1, A) is called
	c := make([]int, n)

	callback(A)

	i := 0
	for i < n {
		if c[i] < i {
			if (i % 2) == 0 {
				A[0], A[i] = A[i], A[0]
			} else {
				A[c[i]], A[i] = A[i], A[c[i]]
			}

			callback(A)

			// Swap has occurred ending the for-loop. Simulate the increment of
			// the for-loop counter
			c[i] += 1
			// Simulate recursive call reaching the base case by bringing the
			// pointer to the base case analog in the array
			i = 0
		} else {
			// Calling generate(i+1, A) has ended as the for-loop terminated.
			// Reset the state and simulate popping the stack by incrementing
			// the pointer.
			c[i] = 0
			i += 1
		}
	}
}

func I64Permutations(A []int64, callback func([]int64)) {
	// Heap's algorithm https://en.wikipedia.org/wiki/Heap%27s_algorithm
	// Using the non-recursive form, which is going to just look wacky since
	// it manually tracks a stack.

	n := int64(len(A))
	// c is an encoding of the stack state. c[k] encodes the for-loop counter
	// for when generate(k+1, A) is called
	c := make([]int64, n)

	callback(A)

	i := int64(0)
	for i < n {
		if c[i] < i {
			if (i % 2) == 0 {
				A[0], A[i] = A[i], A[0]
			} else {
				A[c[i]], A[i] = A[i], A[c[i]]
			}

			callback(A)

			// Swap has occurred ending the for-loop. Simulate the increment of
			// the for-loop counter
			c[i] += 1
			// Simulate recursive call reaching the base case by bringing the
			// point64er to the base case analog in the array
			i = 0
		} else {
			// Calling generate(i+1, A) has ended as the for-loop terminated.
			// Reset the state and simulate popping the stack by incrementing
			// the point64er.
			c[i] = 0
			i += 1
		}
	}
}