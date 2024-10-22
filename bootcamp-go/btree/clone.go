package btree

func (original *BTree) Clone() *BTree {
	if original == nil || original.Root == nil {
		return &BTree{}
	}

	clonedRoot := cloneTree(original.Root)
	clonedTree := &BTree{Root: clonedRoot}
	return clonedTree
}

func cloneTree(node *BTreeNode) *BTreeNode {
	if node == nil {
		return nil
	}

	clonedNode := &BTreeNode{
		Value: node.Value,
	}

	clonedNode.Left = cloneTree(node.Left)
	clonedNode.Right = cloneTree(node.Right)

	if clonedNode.Left != nil {
		clonedNode.Left.Parent = clonedNode
	}
	if clonedNode.Right != nil {
		clonedNode.Right.Parent = clonedNode
	}

	return clonedNode
}
