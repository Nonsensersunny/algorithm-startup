package leetcode

import (
	"testing"
	"time"
)

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{9, 5, 0, 8, 1, 6, 2, 7, 4, 3}
	popped := []int{0, 5, 9, 1, 6, 8, 7, 4, 3, 2}
	now := time.Now()
	res := validateStackSequences(pushed, popped)
	dur := time.Since(now)
	t.Logf("Solved:%v spent:%v", res, dur)
}

func TestMaxSubArray(t *testing.T) {
	data := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	now := time.Now()
	res := maxSubArray(data)
	dur := time.Since(now)
	t.Logf("Solved:%v spent:%v", res, dur)
}
