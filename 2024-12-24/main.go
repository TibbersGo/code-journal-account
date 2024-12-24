package main

func main() {}

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
	}
	for _, n := range m {
		if n != 0 {
			return false
		}
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	m := make(map[int]struct{})
	res := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}

func isHappy(n int) bool {
	m := make(map[int]struct{})
	for n != 1 {
		sum := 0
		for n > 0 {
			tmp := n % 10
			sum += tmp * tmp
			n = n / 10
		}
		if _, ok := m[sum]; ok {
			return false
		} else {
			m[sum] = struct{}{}
		}
		n = sum
	}
	return true
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		} else {
			m[v] = i
		}
	}
	return []int{}
}
