# Tools for leetcode
This package contains useful implementations of data structurs used in leetcode.

# Install
```
> go get github.com/yskang/leetCodeUtil/leetData
```

## treeNode
A treeNode has a one integer value and two pointers which point other treeNode.
```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```
In leetcode problems, the initial value of this structure is provided by a string like this.
```
[3,9,20,null,null,15,7]
```
It is very hard to create test of my own input. So, I make this utils to creating custom testing set easliy.

With this package, you can make TreeNode using a same formatted string in the problem, compare two TreeNode structures, and print TreeNode as a string.

### How to use

Using MakeTreeNode function in this package, you can make new treeNode with a initial string.
```
tree = MakeTreeNode("1,null,2,3,null,4,5")
```
The package has compare and print function. So, you can make a unit test for problems which use treeNode.
```
treeA := MakeTreeNode("1,2,3,4,5")
treeB := MakeTreeNode("1,2,3,4,5")
treeC := MakeTreeNode("1,2,3,4,null,5")
fmt.Println(CompareTreeNode(treeA, treeB))
fmt.Println(CompareTreeNode(treeA, treeC))

// output:
// true
// false

tree = MakeTreeNode("1,null,2,3,null,4,5")
fmt.Println(tree)

// output: 
// 1,null,2,3,null,4,5
```