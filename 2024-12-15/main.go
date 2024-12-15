package main

func main() {}

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其
// 长度。如果不存在符合条件的子数组，返回 0。
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 滑动窗口
	left, sum := 0, 0
	minLen, subLen := 0, 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			subLen = right - left + 1
			if subLen < minLen || minLen == 0 {
				minLen = subLen
			}
			sum -= nums[left]
			left++
		}
	}
	return minLen
}

// 59.
// 给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
func generateMatrix(n int) [][]int {
	// 转几圈
	loop := n / 2
	// 转完后剩下中心点
	mid := n / 2
	// 矩阵元素值
	value := 1
	// 转圈起始位置
	startX, startY := 0, 0
	// 矩阵定义
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 单边收缩长度
	offset := 1
	for loop > 0 {
		x, y := startX, startY
		for y < n-offset {
			matrix[x][y] = value
			value++
			y++
		}
		for x < n-offset {
			matrix[x][y] = value
			value++
			x++
		}
		for y > startY {
			matrix[x][y] = value
			value++
			y--
		}
		for x > startX {
			matrix[x][y] = value
			value++
			x--
		}
		offset++
		startX++
		startY++
		loop--
	}
	if n%2 != 0 {
		matrix[mid][mid] = value
	}
	return matrix
}
