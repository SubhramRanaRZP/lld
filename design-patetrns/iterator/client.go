package main

import (
	"fmt"
	"lld/design-patetrns/iterator/core/data-structures/array-list"
)

func main() {
	dataStructure := array_list.NewArrayList([]int{1, 2, 3})

	fItr := dataStructure.GetForwardIterator()
	for fItr.HasNext() {
		fmt.Print(fItr.GetNext(), " ")
	}
	fmt.Println()

	// again refreshing the forward-iterator
	fItr.Refresh()
	for fItr.HasNext() {
		fmt.Print(fItr.GetNext(), " ")
	}
	fmt.Println()

	// iterating backward
	bItr := dataStructure.GetBackwardIterator()
	for bItr.HasNext() {
		fmt.Print(bItr.GetNext(), " ")
	}
	fmt.Println()
}
