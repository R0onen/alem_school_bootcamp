package bootcamp

import "bootcamp/btree"

func DeleteBtreeLeaves(b *btree.BTree) {
	if b.Root == nil {
		return
	}
	rec3(b.Root)
}

func rec3(b *btree.BTreeNode) int {
	if b == nil {
		return -1
	}
	if b.Left == nil && b.Right == nil {
		return -1
	}
	if rec3(b.Left) < 0 {
		b.Left = nil
	}
	if rec3(b.Right) < 0 {
		b.Right = nil
	}
	return 1
}
