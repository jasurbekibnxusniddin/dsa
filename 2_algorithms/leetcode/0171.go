package leetcode

func titleToNumber(columnTitle string) int {
	var res int
	for _, c := range columnTitle {
		res = res*26 + int(c-'A') + 1
	}
	return res
}
