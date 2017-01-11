package heap

type binaryHeap struct {
	items []int
}

// NewBinaryHeap ceates a new Heap with a binary impmentation
func NewBinaryHeap() Heap {
	return binaryHeap{}
}

func (b binaryHeap) Insert(new int) {
	b.items = append(b.items, new)
	b.rebalance()
}

func (b binaryHeap) Max() int {
	return 0
}

func (b binaryHeap) Peek() int {
	return 0
}

func (b binaryHeap) rebalance() {

}
