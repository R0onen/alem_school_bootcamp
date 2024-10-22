package btree

func (b *BTree) InOrderTraversal(fn func(n *BTreeNode)) {
	b.inOrderTraversalRecursive(b.Root, fn)
}

func (b *BTree) inOrderTraversalRecursive(node *BTreeNode, fn func(n *BTreeNode)) {
	if node == nil {
		return
	}

	b.inOrderTraversalRecursive(node.Left, fn)

	fn(node)

	b.inOrderTraversalRecursive(node.Right, fn)
}
