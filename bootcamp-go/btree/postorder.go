package btree

func (b *BTree) PostOrderTraversal(fn func(n *BTreeNode)) {
	if b.Root == nil {
		return
	}
	recchec2(b.Root, fn)
}

func recchec2(b *BTreeNode, fn func(n *BTreeNode)) {
	if b.Right == nil && b.Left == nil {
		fn(b)
		return
	}
	if b.Left != nil {
		recchec2(b.Left, fn)
	}
	if b.Right != nil {
		recchec2(b.Right, fn)
	}
	fn(b)
}
