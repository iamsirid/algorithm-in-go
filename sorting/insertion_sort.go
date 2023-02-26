package sorting

func InsertionSort(arr []float32) []float32 {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}
