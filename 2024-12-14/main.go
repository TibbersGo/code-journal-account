package main

func main() {}

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// 二分查找, 使用前提，有序+无重复元素
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// left 为左边界，right 为右边界
	// 范围

	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元
// 素的数量。
// 移除元素
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		tmpLeft := nums[left] * nums[left]
		tmpRight := nums[right] * nums[right]
		if tmpLeft <= tmpRight {
			res[i] = tmpRight
			right--
		} else {
			res[i] = tmpLeft
			left++
		}
	}
	return res
}
