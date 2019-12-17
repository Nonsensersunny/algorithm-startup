package leetcode

import "math"

// 946. 验证栈序列
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	var (
		stack = make([]int, len(pushed))
		top   = -1
		popi  = 0
	)
	for i := 0; i < len(pushed); i++ {
		top++
		stack[top] = pushed[i]
		for top >= 0 && stack[top] == popped[popi] {
			stack[top] = 0
			top--
			popi++
		}
	}
	if top != -1 {
		return false
	}
	return true
}

// 53. 最大子序和
func maxSubArray(nums []int) int {
	max, sum := math.MinInt64, 0
	for _, n := range nums {
		if sum > 0 {
			sum += n
			if sum > max {
				max = sum
			}
		} else {
			sum = n
		}
	}
	return max
}
