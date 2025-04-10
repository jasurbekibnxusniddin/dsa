package selectionsort

import "fmt"

func getMin(arr []int) int {

	var minValue = arr[0]
	var minIndex = 0

	for i, e := range arr {

		if e < minValue {
			minValue = e
			minIndex = i
		}
	}

	return minIndex
}

func selectionSort(arr []int) []int {

	var (
		newArr []int
	)

	for len(arr) > 0 {
		minIndex := getMin(arr)
		newArr = append(newArr, arr[minIndex])
		arr = append(arr[:minIndex], arr[minIndex+1:]...)
	}

	return newArr
}

func SelectionSort() {
	fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
}
