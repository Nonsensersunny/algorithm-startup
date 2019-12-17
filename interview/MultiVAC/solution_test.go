package multivac

import (
	"testing"
	"time"
)

func TestFindCommonNumbers(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 7, 9}
	data2 := []int{2, 3, 4, 8, 9}
	now := time.Now()
	res := FindCommonNumbers(data1, data2)
	dur := time.Since(now)
	t.Logf("Common numbers:%v spent:%v", res, dur)
}

func TestBoardShortestPath(t *testing.T) {
	board := [][]int{[]int{3, 0, -2, 4, 0},
		[]int{-1, 2, -2, 1, 4},
		[]int{3, 1, -2, -3, 3},
		[]int{2, -4, -3, -3, 2},
		[]int{5, 2, -2, -3, 1}}
	now := time.Now()
	res := BoardShortestPath(board)
	dur := time.Since(now)
	t.Logf("Min path:%v spent:%v", res, dur)
}

func TestJudgeWinner(t *testing.T) {
	p1 := "4579K"
	p2 := "AAAA2"
	now := time.Now()
	res := JudgeWinner(p1, p2)
	dur := time.Since(now)
	t.Logf("Winner:%v spent:%v", res, dur)
}
