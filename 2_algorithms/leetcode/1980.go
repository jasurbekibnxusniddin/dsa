package leetcode

import "fmt"

func findDifferentBinaryString(nums []string) string {
	var (
		k         = len(nums)
		n         = len(nums[0])
		bins      = make(map[string]struct{})
		backTrack func() bool
		res       = make([]rune, 0, n)
		resp      string
	)

	for _, b := range nums {
		bins[b] = struct{}{}
	}

	backTrack = func() bool {
		if len(res) == n {

			if k < 0 {
				return true
			}
			fmt.Println(string(res), k)
			if _, ok := bins[string(res)]; !ok {
				k--
				resp = string(res)
			}

			return false
		}

		for choice := '0'; choice <= '1'; choice++ {
			res = append(res, choice)
			if backTrack() {
				return true
			}
			res = res[:len(res)-1]
		}

		return false
	}

	backTrack()
	return resp

}

//	func findDifferentBinaryString(nums []string) string {
//		uniqueString := ""
//		for i := 0; i < len(nums[0]); i++ {
//			if nums[i][i] == '0' {
//				uniqueString += "1"
//			} else {
//				uniqueString += "0"
//			}
//		}
//		return uniqueString
//	}
func FindDifferentBinaryString() {

	r := findDifferentBinaryString([]string{"00", "01"})
	fmt.Println(r)
}
