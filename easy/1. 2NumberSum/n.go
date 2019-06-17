package program

import "sort"

// O(n)

func TwoNumberSum(array []int, target int) []int {
	hashTable := map[int]bool{}
	for _, number := range array {
		n := target - number
		if _, ok := hashTable[n]; ok {
			result := []int{number, n}
			sort.Ints(result)
			return result
		} else {
			hashTable[number] = true
		}
	}
	return []int{}
}
