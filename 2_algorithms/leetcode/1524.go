package leetcode

import "fmt"
func numOfSubarrays(arr []int) int {
    mod := int(1e9 + 7)
    evenCount, oddCount := 1, 0 // Prefix sum 0 is even
    prefixSum, result := 0, 0

    for _, num := range arr {
        prefixSum += num
        if prefixSum%2 == 0 {
            result = (result + oddCount) % mod
            evenCount++
        } else {
            result = (result + evenCount) % mod
            oddCount++
        }
    }

    return result
}

func NumOfSubarrays() {
	r := numOfSubarrays([]int{1, 3, 5})
	fmt.Println(r)
}
