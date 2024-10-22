package btree

func (b *BTree) ApplyByLevel(fn func(node *BTreeNode, level int)) {
	if b.Root == nil {
		return
	}
	q := []*BTreeNode{b.Root}
	level := 0

	for len(q) > 0 {
		levelSize := len(q)

		for i := 0; i < levelSize; i++ {
			node := q[i]
			fn(node, level)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		q = q[levelSize:]
		level++
	}
}
