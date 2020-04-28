package tree617

import (
	"testing"
)

var tests = []struct {
	tree1      *TreeNode
	tree2      *TreeNode
	mergedTree *TreeNode
}{
	{
		tree1:      nil,
		tree2:      nil,
		mergedTree: nil,
	},
	{
		tree1:      tree1(),
		tree2:      tree2(),
		mergedTree: tree1PlusTree2(),
	},
	{
		tree1:      tree1(),
		tree2:      nil,
		mergedTree: tree1(),
	},
	{
		tree1:      nil,
		tree2:      tree2(),
		mergedTree: tree2(),
	},
}

func TestMergeTrees(t *testing.T) {
	for _, test := range tests {
		// fmt.Println(traversalTreeNode(test.tree1))
		// fmt.Println(traversalTreeNode(test.tree2))
		// fmt.Println(traversalTreeNode(test.mergedTree))
		result := traversalTreeNode(mergeTrees(test.tree1, test.tree2))
		expectedResult := traversalTreeNode(test.mergedTree)
		if !isEqual(result, expectedResult) {
			t.Errorf("expected %v, got %v\n", expectedResult, result)
		}
	}
}
