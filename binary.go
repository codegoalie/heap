package heap

import "errors"

type BinaryHeap struct {
	items []int
}

// NewBinaryHeap ceates a new Heap with a binary impmentation
func NewBinaryHeap() BinaryHeap {
	return BinaryHeap{}
}

func (b *BinaryHeap) Insert(newb int) {
	b.items = append(b.items, newb)
	b.rebalance()
}

func (b *BinaryHeap) Max() (int, error) {
	if len(b.items) < 1 {
		return 0, errors.New("Heap has no items")
	}
	max := b.items[0]
	b.items = b.items[1:]
	b.rebalance()
	return max, nil
}

func (b BinaryHeap) Peek() (int, error) {
	if len(b.items) < 1 {
		return 0, errors.New("Heap has no items")
	}
	return b.items[0], nil
}

func (b *BinaryHeap) rebalance() {
}
