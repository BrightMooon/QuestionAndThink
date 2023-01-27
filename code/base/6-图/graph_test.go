package graph

import "testing"

func TestXxx(t *testing.T) {

	grid:=make([][]int,8)
	grid[0]=[]int{1,0,0,0,0,0,0,0,0,0,0,0,0}
	grid[1]=[]int{1,0,0,0,0,0,0,0,0,0,0,0,0}
	grid[2]=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0}
	grid[3]=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0}
	grid[4]=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0}
	grid[5]=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0}
	grid[6]=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0}
	grid[7]=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0}
	
		
	maxAreaOfIsland(grid)
}