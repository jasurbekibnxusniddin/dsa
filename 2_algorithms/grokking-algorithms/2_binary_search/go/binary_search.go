package grokkingalgorithms

import "fmt"

func binarySearch(arr []int, target int) *int {

	var (
		low  = 0
		high = len(arr) - 1
	)

	for low <= high {
		
		mid := (low + high) / 2

		if arr[mid] == target {
			return &mid
		}

		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return nil
}

func BinarySearch() {

	r := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9)
	{
		if r == nil {
			fmt.Println("Not found")
		} else {
			fmt.Println(*r)
		}
	}
}
