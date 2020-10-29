package services

import (
	"github.com/Kento75/unittesting-in-golang/golang-testing/src/api/utils/sort"
)

func Sort(elements []int) {
	// 要素数が10000より少ない場合
	if len(elements) <= 10000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
