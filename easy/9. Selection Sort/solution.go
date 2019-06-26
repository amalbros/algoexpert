package program

func SelectionSort(array []int) []int {

	for i := 0; i < len(array)-1; i++ {
		smallest := i
		for j := i + 1; j < len(array); j++ {
			if array[smallest] > array[j] {
				smallest = j
			}
		}
		array[i], array[smallest] = array[smallest], array[i]
	}
	return array
}
