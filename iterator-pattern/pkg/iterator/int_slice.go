package iterator

// IntSliceIterator is an iterator for a slice of integers.
type IntSliceIterator struct {
	slice []int
	index int
}

// NewIntSliceIterator creates a new iterator for the given slice.
func NewIntSliceIterator(slice []int) *IntSliceIterator {
	return &IntSliceIterator{
		slice: slice,
		index: 0,
	}
}

// HasNext checks if there are more elements to iterate over.
func (it *IntSliceIterator) HasNext() bool {
	return it.index < len(it.slice)
}

// Next returns the next element in the iteration and advances the iterator.
func (it *IntSliceIterator) Next() int {
	if !it.HasNext() {
		panic("No more elements in the iterator")
	}
	value := it.slice[it.index]
	it.index++
	return value
}
