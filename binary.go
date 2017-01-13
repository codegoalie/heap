package heap

import "errors"

// BinaryHeap implements a binary tree heap data structure
type BinaryHeap struct {
	items []int
}

// NewBinaryHeap ceates a new Heap with a binary impmentation
func NewBinaryHeap() BinaryHeap {
	return BinaryHeap{}
}

// Insert adds new items to the heap
func (b *BinaryHeap) Insert(newb int) {
	b.items = append(b.items, newb)
	b.rebalance()
}

// Max removes and returns the highest valued item in the heap
func (b *BinaryHeap) Max() (int, error) {
	if len(b.items) < 1 {
		return 0, errors.New("Heap has no items")
	}
	max := b.items[0]
	b.items = b.items[1:]
	b.rebalance()
	return max, nil
}

// Peek returns, but does not remove, the highest valued item in the heap
func (b BinaryHeap) Peek() (int, error) {
	if len(b.items) < 1 {
		return 0, errors.New("Heap has no items")
	}
	return b.items[0], nil
}

func (b *BinaryHeap) rebalance() {
	i := len(b.items) - 1
	for {
		if i < 1 {
			return
		}

		p := parentFor(i)
		if b.items[i] > b.items[p] {
			b.items[i], b.items[p] = b.items[p], b.items[i]
		}
		i = p
	}
}

func parentFor(i int) int {
	if i%2 == 0 {
		return (i / 2) - 1
	}
	return (i - 1) / 2
}
