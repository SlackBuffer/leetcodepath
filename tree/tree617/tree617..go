package tree617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 纯递归
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		t1 = &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}
	if t2 == nil {
		t2 = &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		}
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// peek
func mergeTrees2(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2 // * 全部继承过来
	}
	if t2 == nil {
		return t1 // * 全部继承过来
	}
	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}

// 栈；原位更改 t1 树；递归地先 Right 后 Left
// peak
func mergeTrees3(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	type step [2]*TreeNode
	// var stack []step
	// stack = append(stack, [2]*TreeNode{t1, t2})
	// stack := []step{[2]*TreeNode{t1, t2}}
	stack := []step{step{t1, t2}}
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // * pop

		if t[0] == nil || t[1] == nil {
			continue
		}

		t[0].Val += t[1].Val

		if t[0].Left == nil {
			t[0].Left = t[1].Left // * 全部继承过来，不入栈
		} else {
			stack = append(stack, step{t[0].Left, t[1].Left})
		}

		if t[0].Right == nil {
			t[0].Right = t[1].Right // * 全部继承过来，不入栈
		} else {
			stack = append(stack, step{t[0].Right, t[1].Right})
		}
	}
	return t1
}

// Time complexity : O(m). A total of m nodes need to be traversed. Here, m represents the **minimum** number of nodes from the two given trees.

// Space complexity : O(m). The depth of the recursion tree can go upto m in the case of a skewed tree. In average case, depth will be O(logm).

func isEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func traversalTreeNode(t *TreeNode) []int {
	var s []int
	if t == nil {
		return nil
	}
	s = append(s, t.Val)
	s = append(s, traversalTreeNode(t.Left)...)
	s = append(s, traversalTreeNode(t.Right)...)
	return s
}

func tree1() *TreeNode {
	l3 := newTreeNode(5, nil, nil)
	l21 := newTreeNode(3, l3, nil)
	l22 := newTreeNode(2, nil, nil)
	root := newTreeNode(1, l21, l22)
	return root
}
func tree2() *TreeNode {
	l31 := newTreeNode(4, nil, nil)
	l32 := newTreeNode(7, nil, nil)
	l21 := newTreeNode(1, nil, l31)
	l22 := newTreeNode(3, nil, l32)

	root := newTreeNode(2, l21, l22)
	return root
}
func tree1PlusTree2() *TreeNode {
	l31 := newTreeNode(5, nil, nil)
	l32 := newTreeNode(4, nil, nil)
	l33 := newTreeNode(7, nil, nil)
	l21 := newTreeNode(4, l31, l32)
	l22 := newTreeNode(5, nil, l33)

	root := newTreeNode(3, l21, l22)
	return root
}
func newTreeNode(v int, l, r *TreeNode) *TreeNode {
	t := TreeNode{
		Val:   v,
		Left:  l,
		Right: r,
	}
	return &t
}

// Merge Two Binary Trees (easy)
// Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

// You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

// Example 1:

// ```
// Input:
// 	Tree 1                     Tree 2
//           1                         2
//          / \                       / \
//         3   2                     1   3
//        /                           \   \
//       5                             4   7
// Output:
// Merged tree:
// 	     3
// 	    / \
// 	   4   5
// 	  / \   \
// 	 5   4   7
// ```

// Note: The merging process must start from the root nodes of both trees.
