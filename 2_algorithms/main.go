package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -lstdc++
#include "./leetcode/0171.c"
extern int titleToNumberC(const char* columnTitle);
*/
import "C"
import "github.com/jasurbekibnxusniddin/dsa/algorithms/leetcode/leetcode"

func run() {

	// - leetcode
	// leetcode.GetHappyString()
	// leetcode.FindDifferentBinaryString()
	// leetcode.NumOfSubarrays()
	leetcode.ClosestPrimes()
}

func main() {
	run()
}
