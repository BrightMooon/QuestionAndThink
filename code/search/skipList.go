package search
/*
如何实现一个跳表

*/

const MaxLevel = 7

type Node struct{
	Value int
	Pre *Node
	Next *Node
	Down *Node
}

type SkipList struct{
	Level int
	//每层的头结点
	LevelHeadNode []*Node
}

//新增

//删除

//判断是否存在








