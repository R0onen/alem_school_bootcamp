package bootcamp

import "bootcamp/btree"

func IsBalancedBtree(b *btree.BTree) bool {
	return isBalanced(b.Root)
}

func isBalanced(node *btree.BTreeNode) bool {
	if node == nil {
		return true
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	if abs1(leftHeight-rightHeight) > 1 {
		return false
	}

	return isBalanced(node.Left) && isBalanced(node.Right)
}

func height(node *btree.BTreeNode) int {
	if node == nil {
		return -1 // height of an empty subtree
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func abs1(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
