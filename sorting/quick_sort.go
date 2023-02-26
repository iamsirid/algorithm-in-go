package sorting

func QuickSort(arr []float32) []float32 {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left []float32
	var right []float32

	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)

}
