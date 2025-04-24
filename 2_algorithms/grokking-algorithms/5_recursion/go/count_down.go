package recursion

func countDown(n int) {

	if n < 0 {
		return
	} else {
		println(n)
		countDown(n - 1)
	}
}

func CountDown() {

	countDown(10)
}
