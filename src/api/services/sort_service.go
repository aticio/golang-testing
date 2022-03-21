package services

import "github.com/aticio/golang-testing/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) <= 1000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
