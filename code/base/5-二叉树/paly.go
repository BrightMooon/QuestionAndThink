package main

import "fmt"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main(){
	test:=[]int{1,2,3}
	fmt.Println(test)
	testArray(&test)
	fmt.Println(test)

}


func testArray(array *[]int ){

	*array=append(*array, 8)
}