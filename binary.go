package heap

import "errors"

// BinaryHeap implements a binary tree heap data structure
type BinaryHeap struct {
	items []int
}

// NewBinaryHeap ceates a new Heap with a binary impmentation
func NewBinaryHeap() Heap {
	return &BinaryHeap{}
}

// Insert adds new items to the heap
func (b *BinaryHeap) Insert(newb int) {
	b.items = append(b.items, newb)
	b.upHeap()
}

// Max removes and returns the highest valued item in the heap
func (b *BinaryHeap) Max() (int, error) {
	length := len(b.items)
	if length < 1 {
		return 0, errors.New("Heap has no items")
	}
	max := b.items[0]

	b.items[0] = b.items[length-1]
	b.items = b.items[0 : length-1]
	b.downHeap()
	return max, nil
}

// Peek returns, but does not remove, the highest valued item in the heap
func (b *BinaryHeap) Peek() (int, error) {
	if len(b.items) < 1 {
		return 0, errors.New("Heap has no items")
	}
	return b.items[0], nil
}

func (b *BinaryHeap) upHeap() {
	i := len(b.items) - 1
	for {
		// fmt.Printf("^ Heap(%d): %+v\n", i, b.items)
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

func (b *BinaryHeap) downHeap() {
	b.downHeapFrom(0)
}

func (b *BinaryHeap) downHeapFrom(i int) {
	// fmt.Printf("\\/Heap(%d): %+v\n", i, b.items)
	for _, child := range childrenOf(i) {
		if child < 0 || child >= len(b.items) {
			continue
		}
		if b.items[child] > b.items[i] {
			b.items[i], b.items[child] = b.items[child], b.items[i]
			b.downHeapFrom(child)
		}
	}
}

func childrenOf(i int) [2]int {
	if i == 0 {
		return [2]int{1, 2}
	}
	a := (i + 1) * 2
	b := 2*i + 1
	if i%2 == 0 {
		return [2]int{b, a}
	}
	return [2]int{a, b}
}
