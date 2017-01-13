package heap

import (
	"testing"
)

func TestEmptyMax(t *testing.T) {
	heap := NewBinaryHeap()
	_, err := heap.Max()
	if err == nil {
		t.Errorf("Empty binary heaps should return errors for Max() when empty")
	}
}

func TestMaxForOne(t *testing.T) {
	heap := NewBinaryHeap()
	expected := 4
	heap.Insert(expected)
	max, err := heap.Max()
	if err != nil && max != expected {
		t.Errorf("heap.Max() == %d; expected %d; raised %s", max, expected, err)
	}
}
