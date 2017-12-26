package treeNode

import (
	"fmt"
)

func ExampleMakeTreeNode() {
	node := MakeTreeNode("1,2,3,4,5")
	fmt.Println(node)
	node = MakeTreeNode("1,null,2,3,null,4,5")
	fmt.Println(node)
	node = MakeTreeNode("1")
	node.Left = MakeTreeNode("2")
	node.Right = MakeTreeNode("3")
	node.Left.Left = MakeTreeNode("4")
	node.Right.Right = MakeTreeNode("5")
	fmt.Println(node)
	// output:
	// 1,2,3,4,5
	// 1,null,2,3,null,4,5
	// 1,2,3,4,null,null,5
}

func ExampleCompareTreeNode() {
	nodeA := MakeTreeNode("1,2,3,4,5")
	nodeB := MakeTreeNode("1,2,3,4,5")
	nodeC := MakeTreeNode("1,2,3,4,null,5")
	fmt.Println(CompareTreeNode(nodeA, nodeB))
	fmt.Println(CompareTreeNode(nodeA, nodeC))
	// output:
	// true
	// false
}
