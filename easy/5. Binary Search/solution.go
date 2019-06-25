package program

func BinarySearch(array []int, target int) int {
	// Write your code here.
	return BinaryHelper(array, target, 0, len(array)-1)
}

func BinaryHelper(array []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := int((right + left) / 2)
	element := array[mid]
	if element == target {
		return mid
	} else if target < element {
		return BinaryHelper(array, target, left, mid-1)
	}
	return BinaryHelper(array, target, mid+1, right)
}
