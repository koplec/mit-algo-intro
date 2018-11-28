package main

import "fmt"

func main() {
	ary := []int{5, 2, 4, 6, 1, 3}
	fmt.Printf("before sort -> %v\n", ary)
	sort(ary)
	fmt.Printf("after sort -> %v\n", ary)
}

// insertion-sort
func sort(ary []int) (err error) {
	if len(ary) <= 1 {
		err = nil
		return
	}
	for j := 1; j < len(ary); j++ {
		key := ary[j]
		var i int
		for i = j - 1; i >= 0 && ary[i] > key; i-- {
			ary[i+1] = ary[i] //どんどん右に動かす
		}
		ary[i+1] = key
	}
	return
}
