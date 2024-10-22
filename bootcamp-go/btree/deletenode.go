package btree

func (b *BTree) DeleteNode(value int) {
	if b.Root == nil {
		return
	}

	b.Root = deleteNode(b.Root, value)
}

func deleteNode(root *BTreeNode, value int) *BTreeNode {
	if root == nil {
		return nil
	}

	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		// Case 1: No child or one child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Case 2: Two children
		// Find the minimum node in the right subtree (successor)
		successor := findMin(root.Right)
		root.Value = successor.Value
		root.Right = deleteNode(root.Right, successor.Value)
	}

	return root
}

func findMin(node *BTreeNode) *BTreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
