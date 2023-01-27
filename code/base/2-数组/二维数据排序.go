package array

import (
	"fmt"
	"sort"
)

func Sort2DArray(){
	var intervals = [][]int{{2,3},{1,2},{4,5},{3,4},{9,1}}
	sort.Slice(intervals, func(i,j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Print(intervals)
}
