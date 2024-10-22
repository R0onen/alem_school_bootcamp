package btree

type BTree struct {
	Root *BTreeNode
}

type BTreeNode struct {
	Parent      *BTreeNode
	Left, Right *BTreeNode
	Value       int
}

func NewBTree() *BTree {
	return &BTree{Root: nil}
}

func (b *BTree) ReplaceOrInsert(v int) {
	if b.Root == nil {
		b.Root = &BTreeNode{Parent: nil, Left: nil, Right: nil, Value: v}
		return
	}
	newNode := BTreeNode{Parent: nil, Left: nil, Right: nil, Value: v}
	node := b.Root
	parent := b.Root
	_ = parent
	for node != nil {
		parent = node
		if node.Value > v {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	newNode.Parent = parent
	if parent.Value > v {
		parent.Left = &newNode
	} else {
		parent.Right = &newNode
	}
}
