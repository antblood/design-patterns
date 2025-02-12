package iterator

// IntSet is a simple set implementation for integers using a map.
type IntSet struct {
	elements map[int]struct{}
}

// NewIntSet creates a new empty IntSet.
func NewIntSet() *IntSet {
	return &IntSet{
		elements: make(map[int]struct{}),
	}
}

// Add adds an element to the set.
func (s *IntSet) Add(value int) {
	s.elements[value] = struct{}{}
}

// IntSetIterator is an iterator for the IntSet.
type IntSetIterator struct {
	set    *IntSet
	keys   []int
	cursor int
}

// NewIntSetIterator creates a new iterator for the given set.
func NewIntSetIterator(set *IntSet) *IntSetIterator {
	keys := make([]int, 0, len(set.elements))
	for key := range set.elements {
		keys = append(keys, key)
	}
	return &IntSetIterator{
		set:    set,
		keys:   keys,
		cursor: 0,
	}
}

// HasNext checks if there are more elements to iterate over.
func (it *IntSetIterator) HasNext() bool {
	return it.cursor < len(it.keys)
}

// Next returns the next element in the iteration and advances the iterator.
func (it *IntSetIterator) Next() int {
	if !it.HasNext() {
		panic("No more elements in the iterator")
	}
	value := it.keys[it.cursor]
	it.cursor++
	return value
}
