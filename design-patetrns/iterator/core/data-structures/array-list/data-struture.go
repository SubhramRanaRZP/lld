package array_list

import (
	"lld/design-patetrns/iterator/core"
	"lld/design-patetrns/iterator/core/data-structures/array-list/iterators"
)

// ArrayList is an iterable linear data structure
type ArrayList struct {
	data                              []int
	forwardIterator, backwardIterator core.Iterator
}

func NewArrayList(data []int) *ArrayList {
	return &ArrayList{
		data: data,
	}
}

func (a *ArrayList) GetDataFromCollection(idx int) interface{} {
	if idx >= 0 && idx <= len(a.data) {
		return a.data[idx]
	}
	return nil
}

func (a *ArrayList) GetForwardIterator() core.Iterator {
	if a.forwardIterator == nil {
		a.createForwardIterator()
	}
	return a.forwardIterator
}

func (a *ArrayList) GetBackwardIterator() core.Iterator {
	if a.backwardIterator == nil {
		a.createBackwardIterator()
	}
	return a.backwardIterator
}

func (a *ArrayList) createForwardIterator() {
	a.forwardIterator = iterators.NewForwardIterator(a.data, 0, len(a.data)-1)
}

func (a *ArrayList) createBackwardIterator() {
	a.backwardIterator = iterators.NewBackwardIterator(a.data, 0, len(a.data)-1)
}
