package multivac

import (
	"math"
	"sort"
)

// FindCommonNumbers 现有两个从小到大排好序的 int 数组（每个数组自身没有重复元素）。请找出所有在
// 这两个数组中都出现过的数。请写一个函数，输入为两个数组。
func FindCommonNumbers(data1, data2 []int) []int {
	res := []int{}
	record := map[int]int{}
	for _, i := range data1 {
		record[i]++
	}
	for _, i := range data2 {
		record[i]++
	}
	for k, v := range record {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

// BoardShortestPath 有一个 N 行 M 列的棋盘格，有个国际象棋里的马要从这个棋盘格的第一行
// 跳到最后一行。要求这匹马只能从上往下跳，不能倒着跳，即只能跳往下一行或者下
// 面第二行。
// 每个格子上有一个数字，请为小马寻找一条路径，要求路径上所有数之和最小。
// 小马可以从第一行的任意某个格子开始，也必须到最后一 行
// 的某个格子结束。
// 输入：一个 NxM 的矩阵
// 输出：一个数字，这条路径上所有数之和。
// 附加题思路：如果能够回跳，可以维护一个已访问的bool类型的数组，记录已经到达过的点，然后利用回溯法寻找最短路径
func BoardShortestPath(board [][]int) int {
	// row, col displays numbers of row and column
	row, col := len(board), len(board[0])
	// 违法输入
	if row < 2 || col < 2 || (row+col) < 5 {
		return math.MinInt64
	}
	// 最短路径
	path := math.MaxInt64
	// 当前坐标值
	c, r := 0, 0
	// 目标列坐标值
	ct, rt := 0, row-1
	for ; c < col; c++ {
		for ; ct < col; ct++ {
			ps, sum := hasPath(row, col, r, c, rt, ct), 0
			if ps != nil {
				for _, p := range ps {
					sum += board[p[0]][p[1]]
				}
				if sum < path {
					path = sum
				}
			}
		}
	}
	return path
}

// 判断两点之间是否有路径，并返回路径坐标数组
func hasPath(r, c, r1, c1, r2, c2 int) [][]int {
	// TODO
	return nil
}

// 扑克牌面对应权值
var weights = map[string]int{
	"A": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 10,
	"Q": 10,
	"K": 10,
}

// JudgeWinner “斗牛”是一种热门的扑克游戏，每个人 5 张牌，其中 A 当 1，JQK 都当 10。要求：
// 从 5 张牌中选出 3 张牌（必须刚好 3 张），如果存在 3 张牌加起来是 10 或者 20 或者
// 30，就称为“有牛”，否则就是“没牛”。
// 有牛的情况下，剩余两张牌加起来除以 10 后：
// 1. 余数为 0 就是“牛牛”，是最大的一手牌。
// 2. 如果余数为 9 就是“牛九”，是第二大的，依次类推，“牛一”就是最小。
// 3. “有牛”都大于“没牛”。
// 4. 两人都是没牛的情况下，比最大的一张牌谁大（注意 A 最大，2 最小），最大
// 一张牌一样则比较第二张牌，依次类推。
// 请写一段代码，验证两手牌哪一个最大。
// 输入是长度为 5 的两个字符串，表示两手牌。字符串内容是 23456789TJQKA，注意为
// 了让字符串长度一致，我们用 T 代表 10。
// 请输出一个值：1 表示第一手牌大，-1 表示后一手牌更大，0 表示一样大。
func JudgeWinner(player1, player2 string) int {
	c1, c2 := card2Nums(player1), card2Nums(player2)
	b1, b2 := bull(c1), bull(c2)
	// 都没有牛的情况
	if b1 < 0 && b2 < 0 {
		for i := 0; i < 5; i++ {
			if c1[i] > c2[i] {
				return 1
			}
			if c1[i] < c2[i] {
				return -1
			}
		}
	}
	if b1 > b2 {
		return 1
	}
	if b1 < b2 {
		return -1
	}
	return 0
}

func bull(nums []int) int {
	// 防止有多个不同的牛
	bull := -1
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				s := nums[i] + nums[j] + nums[k]
				if s%10 == 0 {
					curBull := (sum(nums) - s) % 10
					if curBull == 0 {
						// 100 表示牛牛，不用math.MaxInt64防止出现因计算机位数导致不兼容
						return 100
					} else if curBull > bull {
						bull = curBull
					}
				}
			}
		}
	}
	return bull
}

// 求和工具
func sum(data []int) int {
	sum := 0
	for _, d := range data {
		sum += d
	}
	return sum
}

// 将牌面转化为排序的数组
func card2Nums(cards string) []int {
	nums := []int{}
	for _, r := range cards {
		nums = append(nums, weights[string(r)])
	}
	sort.Ints(nums)
	return nums
}
