package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	//var N int
	//fmt.Scanf("%d", &N)
	//a := scanNums(N)
	//a := []int{3, -1, 5}

	left, right, sum := findMaxSubarray(a, 0, len(a)-1)
	fmt.Printf("left=%d right=%d sum=%d\n", left, right, sum)
}

func findMaxSubarray(ary []int, low, high int) (int, int, int) {
	debug("ary=%v low=%v high=%v\n", ary, low, high)
	if low == high {
		return low, high, ary[low]
	}
	mid := (low + high) / 2
	debug(" mid=%v\n", mid)
	leftLow, leftHigh, leftSum := findMaxSubarray(ary, low, mid)
	rightLow, rightHigh, rightSum := findMaxSubarray(ary, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(ary, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}

}

func findMaxCrossingSubarray(ary []int, low, mid, high int) (maxLeft, maxRight, sum int) {
	leftSum := math.MinInt64
	sum = 0
	for i := mid; i >= low; i-- {
		sum = sum + ary[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := math.MinInt64

	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum = sum + ary[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}

	sum = leftSum + rightSum
	return
}

//** DEBUG **/
const DEBUG_ENABLE = false

func debug(format string, a ...interface{}) (n int, err error) {
	if DEBUG_ENABLE {
		return fmt.Fprintf(os.Stdout, format, a...)
	} else {
		return 0, nil
	}
}

func scanNums(len int) (nums []int) {
	nums = make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scan(&nums[i])
	}
	return
}
