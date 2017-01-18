package heap

import (
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

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
	lhs := binomialHeap{firstRoot: &root{node: &node{value: i}}}
	fmt.Printf("Insert(%d)\n", i)

	h.firstRoot = merge(*h, lhs).firstRoot
}

func (h *binomialHeap) Max() (int, error) {
	return h.Peek()
}

func (h *binomialHeap) Peek() (int, error) {
	if h.firstRoot == nil {
		return 0, errors.New("Heap has no items")
	}

	cur := h.firstRoot
	if cur.node == nil {
		return 0, errors.New("Heap's first root has no node")
	}
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

	rhr := rhs.firstRoot
	lhr := lhs.firstRoot

	if rhr == nil {
		return &lhs
	}
	if lhr == nil {
		return &rhs
	}
	if rhr == nil && lhr == nil {
		return &merged
	}

	cur := &root{}
	merged.firstRoot = cur

	for {
		if rhr == nil && lhr == nil {
			break
		}
		fmt.Println("\nIteration")
		fmt.Println("cur")
		spew.Dump(cur)
		fmt.Println("rhr")
		spew.Dump(rhr)
		fmt.Println("lhr")
		spew.Dump(lhr)
		fmt.Println("merged")
		spew.Dump(merged)

		if rhr == nil {
			fmt.Println("No rhr, Keep lhn")
			cur = lhr
			lhr = lhr.next
			cur = cur.next
			if rhr == nil && lhr == nil {
				break
			}
			continue
		}
		if lhr == nil {
			fmt.Println("No lhr, Keep rhn")
			cur = rhr
			rhr = rhr.next
			cur = cur.next
			if rhr == nil && lhr == nil {
				break
			}
			continue
		}

		rhn := rhr.node
		lhn := lhr.node

		if rhn.order == lhn.order {
			fmt.Println("Order equal")
			var carry *node
			if rhn.value > lhn.value {
				carry = &node{order: rhn.order + 1, children: append(rhn.children, lhn), value: rhn.value}
			} else {
				carry = &node{order: lhn.order + 1, children: append(lhn.children, rhn), value: lhn.value}
			}
			cur = &root{node: carry}
			rhr = rhr.next
			lhr = lhr.next
		} else if rhn.order < lhn.order {
			fmt.Println("Keep rhn")
			cur = &root{node: rhn}
			rhr = rhr.next
		} else {
			fmt.Println("Keep lhn")
			cur = &root{node: lhn}
			lhr = lhr.next
		}

		cur = cur.next
	}

	fmt.Println("final merged")
	spew.Dump(merged)
	return &merged
}
