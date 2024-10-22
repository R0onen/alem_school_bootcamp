package btree

func (b *BTree) Get(value int) *BTreeNode {
	node := b.Root

	for node != nil {
		if node.Value == value {
			break
		}
		if node.Value > value {
			node = node.Left
			continue
		}
		node = node.Right
	}

	return node
}
