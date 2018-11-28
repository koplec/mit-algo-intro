package main

import (
	"fmt"
	"math"
)

func main() {
	//a := []int{5, 2}
	a := []int{5, 2, 4, 7, 1, 3, 2, 6}

	fmt.Printf("before sort -> %v\n", a)
	mergeSort(a, 0, len(a)-1)
	fmt.Printf("after sort -> %v\n", a)

}

/**
 * 配列aの中のindex p~rまでをソート
 */
func mergeSort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(a, p, q)
		mergeSort(a, q+1, r)
		merge(a, p, q, r)
	}
}

/**
 * 配列a
 * index p~(q-1), q~rの部分配列をmergeする
 */
func merge(a []int, p, q, r int) {
	n_1 := q - p + 1 //q - (p-1) pからqまでの要素数
	n_2 := r - q     //q+1からrまでの要素数
	left := make([]int, n_1+1)
	right := make([]int, n_2+1)
	for i := 0; i < n_1; i++ {
		left[i] = a[p+i]
	}
	left[n_1] = math.MaxInt32

	for j := 0; j < n_2; j++ {
		right[j] = a[q+j+1]
	}
	right[n_2] = math.MaxInt32

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			a[k] = left[i]
			i = i + 1
		} else { //left[i] > right[j]
			a[k] = right[j]
			j = j + 1
		}
	}
}
