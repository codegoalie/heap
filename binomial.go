package heap

import "errors"

type root struct {
	node *node
	next *root
}

type node struct {
	value    int
	order    int
	children []*node
}

type binomialHeap struct {
	firstRoot *root
}

// NewBinomialHeap creates a new heap with a binomal tree based implementation
func NewBinomialHeap() Heap {
	return &binomialHeap{}
}

func (h *binomialHeap) Insert(i int) {
	new := binomialHeap{firstRoot: &root{node: &node{value: i}}}

	h = merge(*h, new)
}

func (h *binomialHeap) Max() (int, error) {
	panic("not implemented")
}

func (h *binomialHeap) Peek() (int, error) {
	if h.firstRoot == nil {
		return 0, errors.New("Heap has no items")
	}

	cur := h.firstRoot
	max := cur.node.value

	for {
		if cur.next == nil {
			break
		}

		cur = cur.next
		if max < cur.node.value {
			max = cur.node.value
		}
	}

	return max, nil
}

func merge(rhs, lhs binomialHeap) *binomialHeap {
	merged := binomialHeap{}
	order := 0

	for {

	}

	return &merged
}
