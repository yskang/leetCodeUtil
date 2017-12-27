# Tools for leetcode
This package contains useful implementations of data structurs used in leetcode.

## treeNode
A treeNode has a one integer value and two pointers which pointing other treeNode.
```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```
In leetcode problem, the initial value of this structure is provided by string like this.
```
[3,9,20,null,null,15,7]
```
Using MakeTreeNode function, you can make new treeNode with a initial string.
```
node = MakeTreeNode("1,null,2,3,null,4,5")
```
The package has compare and print function. So, you can make a unit test for problems which use treeNode.
```
nodeA := MakeTreeNode("1,2,3,4,5")
nodeB := MakeTreeNode("1,2,3,4,5")
nodeC := MakeTreeNode("1,2,3,4,null,5")
fmt.Println(CompareTreeNode(nodeA, nodeB))
fmt.Println(CompareTreeNode(nodeA, nodeC))

// output:
// true
// false

node = MakeTreeNode("1,null,2,3,null,4,5")
fmt.Println(node)

// output: 
// 1,null,2,3,null,4,5
```