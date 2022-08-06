package iterators

import (
	"lld/design-patetrns/iterator/core"
)

type forwardIterator struct {
	startIndex, endIndex, idx int
	data                      []int
}

func NewForwardIterator(data []int, startIndex, endIndex int) core.Iterator {
	return &forwardIterator{
		startIndex: startIndex,
		endIndex:   endIndex,
		idx:        startIndex - 1,
		data:       data,
	}
}

func (i *forwardIterator) HasNext() bool {
	return i.idx < i.endIndex
}

func (i *forwardIterator) GetNext() interface{} {
	if i.HasNext() {
		i.idx++
		return i.data[i.idx]
	}
	return nil
}

func (i *forwardIterator) Refresh() {
	i.idx = i.startIndex - 1
}
