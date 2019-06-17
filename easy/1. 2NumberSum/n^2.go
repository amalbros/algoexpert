package program

import "sort"

// O(n^2)

func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return_Var := []int{array[i], array[j]}
				sort.Ints(return_Var)
				return return_Var
			}
		}
	}
	return []int{}
}
