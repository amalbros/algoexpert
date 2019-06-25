package program

import "math"

func FindThreeLargestNumbers(array []int) []int {
	the_three := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, num := range array {
		the_three = update(the_three, num)
	}
	return the_three
}

func update(array []int, num int) []int {
	if num > array[2] {
		return shift_numbers(array, num, 2)
	} else if num > array[1] {
		return shift_numbers(array, num, 1)
	} else if num > array[0] {
		return shift_numbers(array, num, 0)
	}
	return array
}

func shift_numbers(array []int, num int, index int) []int {
	for i := 0; i <= index; i++ {
		if i == index {
			array[i] = num
		} else {
			array[i] = array[i+1]
		}
	}
	return array
}
