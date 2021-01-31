// Triple Step: A child is running up a staircase with n steps and

package p8_1

func TripleStep_Recursion(steps_left uint32) uint32 {
	if steps_left <= 3 {
		return steps_left
	}

	return TripleStep_Recursion(steps_left-3) + TripleStep_Recursion(steps_left-2) + TripleStep_Recursion(steps_left-1)
}

func Get_TripleStep_Memo() (implFunc func(uint32) uint32) {
	memo := make(map[uint32]uint32)

	implFunc = func(steps_left uint32) uint32 {
		if steps_left <= 3 {
			return steps_left
		} else if result, ok := memo[steps_left]; ok {
			return result
		}

		memo[steps_left] = implFunc(steps_left-1) + implFunc(steps_left-2) + implFunc(steps_left-3)
		return memo[steps_left]
	}

	return implFunc
}
