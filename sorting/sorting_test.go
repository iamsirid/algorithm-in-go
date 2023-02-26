package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestSuite struct {
	input    []float32
	expected []float32
}

func getTestCases() []TestSuite {
	return []TestSuite{
		{[]float32{}, []float32{}},
		{[]float32{5}, []float32{5}},
		{[]float32{1, 2, 3, 4, 5}, []float32{1, 2, 3, 4, 5}},
		{[]float32{5, 4, 3, 2, 1}, []float32{1, 2, 3, 4, 5}},
		{[]float32{4, 8, 2, 1, 5, 3, 6, 7}, []float32{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]float32{-3, 5, 1, 0, -2, 4, -1}, []float32{-3, -2, -1, 0, 1, 4, 5}},
		{[]float32{3.14, 1.618, 2.718, 0.707, 1.414}, []float32{0.707, 1.414, 1.618, 2.718, 3.14}},
	}
}

func runningSortingTest(t *testing.T, sortingFunc func([]float32) []float32) {
	tests := getTestCases()

	for _, tc := range tests {
		result := sortingFunc(tc.input)
		assert.Equal(t, tc.expected, result, "They should be equal")
	}
}
func TestInsertionSort(t *testing.T) {
	runningSortingTest(t, InsertionSort)
}

func TestBubbleSort(t *testing.T) {
	runningSortingTest(t, BubbleSort)
}

func TestSelectionSort(t *testing.T) {
	runningSortingTest(t, SelectionSort)
}

func TestQuickSort(t *testing.T) {
	runningSortingTest(t, QuickSort)
}

func TestMergeSort(t *testing.T) {
	runningSortingTest(t, MergeSort)
}
