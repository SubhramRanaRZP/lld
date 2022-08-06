package core

// Iterator is a common interface which is reusable for any kind of iterableDataStructure
// to iterate any kind of traversal. Just need to implement in that way
/*
1. Create a new iterable data structure under 'data-structures' folder similar to 'array-list'
2. If you want to create a new kind of iterator for an existing dat structure, then
	add a new iterator under the same data structure and implement Iterator interface.
	ex: similar to 'data-structures/array-list/iterators/forward-iterator.go'
*/
type Iterator interface {
	HasNext() bool
	GetNext() interface{}
	Refresh()
}
