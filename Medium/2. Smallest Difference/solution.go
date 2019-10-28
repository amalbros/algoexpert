package program

import (
	"math"
	"sort"
)

func SmallestDifference(array1 []int, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	index1, index2 := 0, 0
	small, current := math.MaxInt32, math.MaxInt32
	small_pair := []int{}
	for index1 < len(array1) && index2 < len(array2) {
		first, second := array1[index1], array2[index2]
		if first < second {
			current = second - first
			index1 += 1
		} else if second < first {
			current = first - second
			index2 += 1
		} else {
			return []int{first, second}
		}
		if small > current {
			small = current
			small_pair = []int{first, second}
		}
	}
	return small_pair

}
