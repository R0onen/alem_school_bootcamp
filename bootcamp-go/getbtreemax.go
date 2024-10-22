package bootcamp

import "bootcamp/btree"

func GetMax(b *btree.BTree) *btree.BTreeNode {
	node := b.Root
	for node.Right != nil {
		node = node.Right
	}

	return node
}
