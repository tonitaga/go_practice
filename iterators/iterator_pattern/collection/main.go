package main

import "fmt"

type CollectionItem struct {
	Field1 int
	Field2 int
}

type SomeCollection struct {
	Items []*CollectionItem
}

type SomeCollectionIterator struct {
	index int
	items []*CollectionItem
}

func (it *SomeCollectionIterator) HasNext() bool {
	return it.index < len(it.items)
}

func (it *SomeCollectionIterator) GetNext() *CollectionItem {
	item := it.items[it.index]
	it.index++
	return item
}

func (s *SomeCollection) GetIterator() *SomeCollectionIterator {
	return &SomeCollectionIterator{
		index: 0,
		items: s.Items,
	}
}

func main() {
	someCollection := &SomeCollection{
		Items: []*CollectionItem{
			{
				Field1: 0,
				Field2: 1,
			},
			{
				Field1: 2,
				Field2: 3,
			},
		},
	}

	iterator := someCollection.GetIterator()
	for iterator.HasNext() {
		item := iterator.GetNext()
		fmt.Println(*item)
	}
}
