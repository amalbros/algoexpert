package program

import "sort"

// O(log(n))

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			result := []int{array[left], array[right]}
			return result
		} else if sum < target {
			left += 1
		} else if sum > target {
			right -= 1
		}
	}
	return []int{}
}
