package btree

func (b *BTree) PreOrderTraversal(fn func(n *BTreeNode)) {
	b.preOrderTraversalRecursive(b.Root, fn)
}

func (b *BTree) preOrderTraversalRecursive(node *BTreeNode, fn func(n *BTreeNode)) {
	if node == nil {
		return
	}

	fn(node)

	b.preOrderTraversalRecursive(node.Left, fn)

	b.preOrderTraversalRecursive(node.Right, fn)
}
