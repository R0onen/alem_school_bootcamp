package bootcamp

import "bootcamp/btree"

func GetMin(b *btree.BTree) *btree.BTreeNode {
	node := b.Root
	for node.Left != nil {
		node = node.Left
	}

	return node
}
