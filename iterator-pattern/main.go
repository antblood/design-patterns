package main

import (
	"fmt"

	"github.com/antblood/design-patterns/iterator-pattern/pkg/iterator"
)

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Create an iterator for the slice
	sliceIterator := iterator.NewIntSliceIterator(numbers)

	// Iterate over the elements
	for sliceIterator.HasNext() {
		number := sliceIterator.Next()
		fmt.Println(number)
	}

	// Create a new IntSet and add elements to it
	set := iterator.NewIntSet()
	set.Add(10)
	set.Add(20)
	set.Add(30)

	// Create an iterator for the set
	setIterator := iterator.NewIntSetIterator(set)

	// Iterate over the elements in the set
	fmt.Println("Iterating over set:")
	for setIterator.HasNext() {
		fmt.Println(setIterator.Next())
	}
}
