package heap

type Heap interface {
	Insert(int)
	Max() (int, error)
	Peek() (int, error)
}
