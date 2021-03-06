package heap

import (
	"testing"
)

func TestBinaryEmptyMax(t *testing.T) {
	heap := NewBinaryHeap()
	_, err := heap.Max()
	if err == nil {
		t.Errorf("Empty binary heaps should return errors for Max() when empty")
	}
}

func TestBinaryMaxForOne(t *testing.T) {
	heap := NewBinaryHeap()
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

func TestBinaryMaxForTwo(t *testing.T) {
	heap := NewBinaryHeap()
	expected := 4
	heap.Insert(3)
	heap.Insert(expected)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != expected {
		t.Errorf("heap.Max() == %d; expected %d", max, expected)
	}
}

func TestBinaryMaxForThree(t *testing.T) {
	heap := NewBinaryHeap()
	expected := 4
	heap.Insert(3)
	heap.Insert(1)
	heap.Insert(expected)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != expected {
		t.Errorf("heap.Max() == %d; expected %d", max, expected)
	}
}

func TestBinaryMaxForTwenty(t *testing.T) {
	heap := NewBinaryHeap()
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

func TestBinaryMaxRemoves(t *testing.T) {
	heap := NewBinaryHeap()
	highest := 8
	next := 7
	third := 6
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(third)   // 6
	heap.Insert(next)    // 7
	heap.Insert(highest) // 8

	// 2
	// 2 3
	// 3 2
	// 3 2 4
	// 4 2 3
	// 4 2 3 5
	// 4 5 3 2
	// 5 4 3 2
	// 5 4 3 2 6
	// 5 6 3 2 4
	// 6 5 3 2 4
	// 6 5 3 2 4 7
	// 6 5 7 2 4 3
	// 7 5 6 2 4 3
	// 7 5 6 2 4 3 8
	// 7 5 8 2 4 3 6
	// 8 5 7 2 4 3 6
	// 6 5 7 2 4 3

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

func TestBinaryPeekNoRemove(t *testing.T) {
	heap := NewBinaryHeap()
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

func TestBinaryParentFor(t *testing.T) {
	cases := map[int]int{
		1:  0,
		2:  0,
		3:  1,
		4:  1,
		5:  2,
		6:  2,
		7:  3,
		8:  3,
		9:  4,
		10: 4,
		11: 5,
		12: 5,
	}

	for i, expected := range cases {
		actual := parentFor(i)
		if actual != expected {
			t.Errorf("parentFor(%d) == %d; expected %d", i, actual, expected)
		}
	}
}
