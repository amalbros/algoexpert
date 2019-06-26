package program

func BubbleSort(array []int) []int {
	sorted := false
	count := 0

	for !sorted {
		sorted = true
		for i := 0; i < len(array)-count-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
		count++
	}
	return array
}
