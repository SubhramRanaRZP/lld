package iterators

import (
	"lld/design-patetrns/iterator/core"
)

type backwardIterator struct {
	startIndex, endIndex, idx int
	data                      []int
}

func NewBackwardIterator(data []int, startIndex, endIndex int) core.Iterator {
	return &backwardIterator{
		startIndex: startIndex,
		endIndex:   endIndex,
		idx:        endIndex + 1,
		data:       data,
	}
}

func (i *backwardIterator) HasNext() bool {
	return i.idx > i.startIndex
}

func (i *backwardIterator) GetNext() interface{} {
	if i.HasNext() {
		i.idx--
		return i.data[i.idx]
	}
	return nil
}

func (i *backwardIterator) Refresh() {
	i.idx = i.endIndex + 1
}
