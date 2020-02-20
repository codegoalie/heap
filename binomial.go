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
	max, err := h.findMax()
	if err != nil {
		return 0, err
	}

	return max.value, nil
}

func (h *binomialHeap) Peek() (int, error) {
	max, err := h.findMax()
	if err != nil {
		return 0, err
	}
	return max.value, nil
}

func (h *binomialHeap) findMax() (*node, error) {
	if h.firstRoot == nil {
		return nil, errors.New("Heap has no items")
	}

	cur := h.firstRoot
	if cur.node == nil {
		return nil, errors.New("Heap's first root has no node")
	}
	max := cur.node

	for {
		if cur.next == nil {
			break
		}

		cur = cur.next
		if max.value < cur.node.value {
			max = cur.node
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

	next := &root{}
	cur := &root{}
	merged.firstRoot = cur

	for {
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
			next.node = lhr.node
			lhr = lhr.next
		} else if lhr == nil {
			fmt.Println("No lhr, Keep rhn")
			next.node = rhr.node
			rhr = rhr.next
		} else {
			rhn := rhr.node
			lhn := lhr.node

			if rhn.order == lhn.order {
				fmt.Println("Order equal")
				merged, _ := mergeTrees(*rhn, *lhn)
				next.node = merged
				rhr = rhr.next
				lhr = lhr.next
			} else if rhn.order < lhn.order {
				fmt.Println("Keep rhn")
				next.node = rhn
				rhr = rhr.next
			} else {
				fmt.Println("Keep lhn")
				next.node = lhn
				lhr = lhr.next
			}
		}

		if merged.firstRoot.node == nil {
			merged.firstRoot = next
		} else {
			cur.next = next
		}

		if rhr == nil && lhr == nil {
			break
		}

		cur = next
		next = &root{}

	}
	fmt.Println("last cur")
	spew.Dump(cur)

	fmt.Println("final merged")
	spew.Dump(merged)
	return &merged
}

func mergeTrees(first, second node) (*node, error) {
	if first.order != second.order {
		return &node{}, fmt.Errorf("Trees must have the same order to be merged. %v != %v", first, second)
	}

	merged := node{order: first.order + 1}
	if first.value > second.value {
		merged.value = first.value
		merged.children = []*node{&second}
	} else {
		merged.value = second.value
		merged.children = []*node{&first}
	}
	return &merged, nil
}
