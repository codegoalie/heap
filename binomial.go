package heap

import (
	"errors"
	"fmt"
)

type root struct {
	node     *node
	next     *root
	previous *root
}

type node struct {
	value    int
	children []*node
}

func (n *node) order() int {
	return len(n.children)
}

type binomialHeap struct {
	firstRoot *root
}

// NewBinomialHeap creates a new heap with a binomal tree based implementation
func NewBinomialHeap() Heap {
	return &binomialHeap{}
}

func (h *binomialHeap) Insert(i int) {
	newFirst := &root{node: &node{value: i}, next: h.firstRoot}
	if h.firstRoot != nil {
		h.firstRoot.previous = newFirst
	}
	newFirst.next = h.firstRoot
	h.firstRoot = newFirst
	h.combine()
}

func (h *binomialHeap) Max() (int, error) {
	max, err := h.findMax()
	if err != nil {
		return 0, err
	}

	maxValue := max.node.value
	if max.previous != nil {
		max.previous.next = max.next
	} else {
		h.firstRoot = max.next
	}

	if max.next != nil {
		max.next.previous = max.previous
	}

	for _, tree := range max.node.children {
		newRoot := &root{node: tree}
		h.insertRoot(newRoot)
	}

	return maxValue, nil
}

func (h *binomialHeap) Peek() (int, error) {
	max, err := h.findMax()
	if err != nil {
		return 0, err
	}
	return max.node.value, nil
}

func (h *binomialHeap) insertRoot(newRoot *root) {
	curRoot := h.firstRoot
	if curRoot == nil {
		h.firstRoot = newRoot
		return
	}

	newOrder := newRoot.node.order()

	for {
		if newOrder <= curRoot.node.order() {
			newRoot.next = curRoot
			if curRoot.previous != nil {
				curRoot.previous.next = newRoot
				newRoot.previous = curRoot.previous
			} else {
				h.firstRoot = newRoot
			}
			curRoot.previous = newRoot

			break
		}

		if curRoot.next == nil {
			newRoot.previous = curRoot
			curRoot.next = newRoot
			break
		}
		curRoot = curRoot.next
	}
	h.combine()
}

func (h *binomialHeap) findMax() (*root, error) {
	if h.firstRoot == nil {
		return nil, errors.New("Heap has no items")
	}

	cur := h.firstRoot
	if cur.node == nil {
		return nil, errors.New("Heap's first root has no node")
	}
	max := cur

	for {
		if cur.next == nil {
			break
		}

		cur = cur.next
		if max.node.value < cur.node.value {
			max = cur
		}
	}

	return max, nil
}

func (h *binomialHeap) combine() {
	curRoot := h.firstRoot
	for {
		if curRoot == nil || curRoot.next == nil {
			return
		}

		if curRoot.node.order() != curRoot.next.node.order() {
			curRoot = curRoot.next
			continue
		}

		newNode, err := mergeTrees(curRoot.node, curRoot.next.node)
		if err != nil {
			err = fmt.Errorf("failed to combine same order roots: %w", err)
			fmt.Println(err)
			return
		}

		curRoot.node = newNode
		nextNext := curRoot.next.next
		curRoot.next = nextNext
		if nextNext != nil {
			nextNext.previous = curRoot
		}
	}
}

func mergeTrees(first, second *node) (*node, error) {
	if first.order() != second.order() {
		return &node{}, fmt.Errorf("Trees must have the same order to be merged. %v != %v", first, second)
	}

	if first.value > second.value {
		first.children = append(first.children, second)
		return first, nil
	}

	second.children = append(second.children, first)
	return second, nil
}
