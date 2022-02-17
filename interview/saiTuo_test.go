package interview

import (
	"fmt"
	"testing"
)

type Tree struct {
	val   int
	Right *Tree
	Left  *Tree
}

var Result = []*Tree{}

func TestSai(t *testing.T) {
	// 树的层次输出
	trees := Tree{val: 1}
	trees.Left = &Tree{val: 2}
	trees.Right = &Tree{val: 3}
	trees.Left.Left = &Tree{val: 4}
	trees.Left.Right = &Tree{val: 5}
	trees.Right.Left = &Tree{val: 6}
	trees.Right.Right = &Tree{val: 7}

	treeLayerPrint(&trees)
	for _, item := range Result {
		fmt.Println(&item.val)
	}
}

func treeLayerPrint(tree *Tree) {
	Result = append(Result, tree)
	if &(tree.Left) != nil {
		Result = append(Result, tree.Left)
	}
	if &(tree.Right) != nil {
		Result = append(Result, tree.Left)
	}
	treeLayerPrint(tree.Left)
	treeLayerPrint(tree.Right)
}
