package sort

import "fmt"

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	//把一个数组一分为二
	mid := length / 2
	left := arr[:mid]
	right := arr[mid:]
	fmt.Printf("left=%v,right=%v \n",left,right)
	return merge(mergeSort(left),mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}