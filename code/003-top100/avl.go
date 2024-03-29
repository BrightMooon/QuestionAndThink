package main
import (
	"fmt"
)


type avlTreeNode struct {
	key   int
	high  int
	left  *avlTreeNode
	right *avlTreeNode
}

func NewAVLTreeNode(value int) *avlTreeNode{
	return &avlTreeNode{key:value}
}


func highTree(p *avlTreeNode) int {
	if p == nil {
		return -1
	} else {
		return p.high
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// look LL
func left_left_rotation(k *avlTreeNode) *avlTreeNode {
	var kl *avlTreeNode
	kl = k.left
	k.left = kl.right
	kl.right = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kl.high = max(highTree(kl.left), k.high) + 1
	return kl
}

//look RR
func right_right_rotation(k *avlTreeNode) *avlTreeNode {
	var kr *avlTreeNode
	kr = k.right
	k.right = kr.left
	kr.left = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kr.high = max(k.high, highTree(kr.right)) + 1
	return kr
}


func left_righ_rotation(k *avlTreeNode) *avlTreeNode {
	k.left = right_right_rotation(k.left)
	return left_left_rotation(k)
}

func right_left_rotation(k *avlTreeNode) *avlTreeNode {
	k.right = left_left_rotation(k.right)
	return right_right_rotation(k)
}

//insert to avl
func avl_insert(avl *avlTreeNode, key int) *avlTreeNode {
	if avl == nil {
		avl = NewAVLTreeNode(key)
	} else if key < avl.key {
		avl.left = avl_insert(avl.left, key)
		if highTree(avl.left)-highTree(avl.right) == 2 {
			if key < avl.left.key { //LL
				avl = left_left_rotation(avl)
			} else { // LR
				avl = left_righ_rotation(avl)
			}
		}
	} else if key > avl.key {
		avl.right = avl_insert(avl.right, key)
		if (highTree(avl.right) - highTree(avl.left)) == 2 {
			if key < avl.right.key { // RL
				avl = right_left_rotation(avl)
			} else {
				fmt.Println("right right", key)
				avl = right_right_rotation(avl)
			}
		}
	} else if key == avl.key {
		fmt.Println("the key", key, "has existed!")
	}
	//notice: update high(may be this insert no rotation, so you should update high)
	avl.high = max(highTree(avl.left), highTree(avl.right)) + 1
	return avl
}


//display avl tree  key by asc
func displayAsc(avl *avlTreeNode) []int{
	return appendValues([]int{},avl)
}

func appendValues(values []int, avl *avlTreeNode)  []int{
	if avl != nil {
		values = appendValues(values,avl.left)
		values = append(values,avl.key)
		values = appendValues(values,avl.right)
	}
	return values
}

//display avl tree key by desc
func displayDesc(avl *avlTreeNode) []int {
	return appendValues2([]int{},avl)
}

func appendValues2(values []int, avl *avlTreeNode) []int {
	if avl != nil {
		values = appendValues2(values,avl.right)
		values = append(values,avl.key)
		values = appendValues2(values,avl.left)
	}
	return values
}




