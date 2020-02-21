package heap

import (
	"testing"
)

func TestBinomialTreeMergeTwoZeros(t *testing.T) {
	highestValue := 6
	first := &node{value: 4}
	second := &node{value: highestValue}
	merged, err := mergeTrees(first, second)
	if err != nil {
		t.Errorf("mergeTrees(%v, %v) raised %s, but shouldn't", first, second, err.Error())
	}

	if merged.order() != 1 {
		t.Errorf("mergeTrees(%v, %v) should have order of 1, but has %d", first, second, merged.order())
	}

	if merged.value != highestValue {
		t.Errorf("mergeTrees(%v, %v) should have value of %d, but has %d", first, second, merged.value, highestValue)
	}
}

func TestBinomialTreeMergeTwoOnes(t *testing.T) {
	highestValue := 6
	first := &node{value: highestValue, children: []*node{{value: 2}}}
	second := &node{value: 4, children: []*node{{value: 3}}}
	merged, err := mergeTrees(first, second)
	if err != nil {
		t.Errorf("mergeTrees(%v, %v) raised %s, but shouldn't", first, second, err.Error())
	}

	if merged.order() != 2 {
		t.Errorf("mergeTrees(%v, %v) should have order of 1, but has %d", first, second, merged.order())
	}

	if merged.value != highestValue {
		t.Errorf("mergeTrees(%v, %v) should have value of %d, but has %d", first, second, merged.value, highestValue)
	}
}

func TestBinomialTreeMergeMismatch(t *testing.T) {
	first := &node{value: 4}
	second := &node{value: 5, children: []*node{{value: 3}}}
	merged, err := mergeTrees(first, second)
	if err == nil {
		t.Errorf("mergeTrees(%v, %v) should have returned an error, but didn't", first, second)
	}

	if merged.order() != 0 {
		t.Errorf("mergeTrees(%v, %v) should have a zero-value order, but has %d", first, second, merged.order())
	}

	if merged.value != 0 {
		t.Errorf("mergeTrees(%v, %v) should have zero-value value, but has %d", first, second, merged.value)
	}
}

func TestBinomialEmptyMax(t *testing.T) {
	heap := NewBinomialHeap()
	_, err := heap.Max()
	if err == nil {
		t.Errorf("Empty binary heaps should return errors for Max() when empty")
	}
}

func TestBinomialMaxForOne(t *testing.T) {
	heap := NewBinomialHeap()
	expected := 4
	heap.Insert(expected)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != expected {
		t.Errorf("heap.Max() == %d; expected %d", max, expected)
	}
}

func TestBinomialMaxForTwo(t *testing.T) {
	heap := NewBinomialHeap()
	expected := 4
	heap.Insert(expected)
	heap.Insert(3)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != expected {
		t.Errorf("heap.Max() == %d; expected %d", max, expected)
	}
}

func TestBinomialMaxForThree(t *testing.T) {
	heap := NewBinomialHeap()
	expected := 4
	heap.Insert(3)
	heap.Insert(expected)
	heap.Insert(1)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != expected {
		t.Errorf("heap.Max() == %d; expected %d", max, expected)
	}
}

func TestBinomialMaxForTwenty(t *testing.T) {
	heap := NewBinomialHeap()
	for i := 1; i < 21; i++ {
		heap.Insert(i)
	}
	expected := 21
	heap.Insert(expected)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != expected {
		t.Errorf("heap.Max() == %d; expected %d", max, expected)
	}
}

func TestBinomialMaxRemoves(t *testing.T) {
	heap := NewBinomialHeap()
	highest := 8
	next := 7
	third := 6
	heap.Insert(third) // 6
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(next)    // 7
	heap.Insert(highest) // 8
	heap.Insert(4)
	heap.Insert(5)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != highest {
		t.Errorf("heap.Max() == %d; expected %d", max, highest)
	}

	nextMax, err := heap.Max()
	if err != nil {
		t.Errorf("2nd heap.Max() raised %s; expected no error", err.Error())
	}
	if nextMax != next {
		t.Errorf("2nd heap.Max() == %d; expected %d", nextMax, next)
	}

	thirdMax, err := heap.Max()
	if err != nil {
		t.Errorf("3nd heap.Max() raised %s; expected no error", err.Error())
	}
	if thirdMax != third {
		t.Errorf("3nd heap.Max() == %d; expected %d", thirdMax, third)
	}
}

func TestBinomialPeekNoRemove(t *testing.T) {
	heap := NewBinomialHeap()
	highest := 4
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(highest)

	peek, err := heap.Peek()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if peek != highest {
		t.Errorf("heap.Max() == %d; expected %d", peek, highest)
	}

	nextPeek, err := heap.Peek()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if nextPeek != highest {
		t.Errorf("heap.Max() == %d; expected %d", nextPeek, highest)
	}
}
