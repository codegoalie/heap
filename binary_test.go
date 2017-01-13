package heap

import "testing"

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

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != expected {
		t.Errorf("heap.Max() == %d; expected %d", max, expected)
	}
}

func TestMaxForTwo(t *testing.T) {
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

func TestMaxForThree(t *testing.T) {
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

func TextMaxForTwenty(t *testing.T) {
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

func TestMaxRemoves(t *testing.T) {
	heap := NewBinaryHeap()
	highest := 4
	next := 3
	heap.Insert(2)
	heap.Insert(next)
	heap.Insert(highest)

	max, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if max != highest {
		t.Errorf("heap.Max() == %d; expected %d", max, highest)
	}

	nextMax, err := heap.Max()

	if err != nil {
		t.Errorf("heap.Max() raised %s; expected no error", err.Error())
	}
	if nextMax != next {
		t.Errorf("heap.Max() == %d; expected %d", nextMax, next)
	}
}

func TestPeekNoRemove(t *testing.T) {
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

func TestParentFor(t *testing.T) {
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
