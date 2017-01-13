package heap

// Heap defines how a user will use a heap object
type Heap interface {
	Insert(int)
	Max() (int, error)
	Peek() (int, error)
}
