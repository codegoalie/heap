package heap

type Heap interface {
	Insert(int)
	Max() int
	Peek() int
}
