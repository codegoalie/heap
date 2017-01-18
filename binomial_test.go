package heap

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

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
	spew.Dump("Insert 4: %+v\n", heap)
	heap.Insert(3)
	spew.Dump("Insert 3: %+v\n", heap)

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
	fmt.Println("\npost-inserts")
	spew.Dump(heap)

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

func TestBinomialMergeZeroEmpty(t *testing.T) {
	expectedValue := 4
	one := binomialHeap{firstRoot: &root{node: &node{value: expectedValue}}}
	zero := binomialHeap{}
	merged := merge(one, zero)
	if merged.firstRoot == nil {
		t.Error("Expected merge of zero and empty heaps to have one root, but none")
	} else {
		if merged.firstRoot.next != nil {
			t.Error("Expected merge of zero and empty heaps to have one root, but has more")
		}

		if merged.firstRoot.node.value != expectedValue {
			t.Errorf("Expected merge of zero and empty heaps to have value of %d, but is %d", expectedValue, merged.firstRoot.node.value)
		}
	}
}

func TestBinomialMergeEmptyZero(t *testing.T) {
	expectedValue := 4
	one := binomialHeap{firstRoot: &root{node: &node{value: expectedValue}}}
	zero := binomialHeap{}
	merged := merge(zero, one)
	if merged.firstRoot == nil {
		t.Error("Expected merge of zero and empty heaps to have one root, but none")
	} else {
		if merged.firstRoot.next != nil {
			t.Error("Expected merge of zero and empty heaps to have one root, but has more")
		}

		if merged.firstRoot.node.value != expectedValue {
			t.Errorf("Expected merge of zero and empty heaps to have value of %d, but is %d", expectedValue, merged.firstRoot.node.value)
		}
	}
}
