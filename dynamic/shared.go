package dynamic

// Iterater is an interface that defines methods for an iterator over a
// collection of elements of type T.
type Iterater[T any] interface {
	// Consume advances the iterator to the next element in the
	// collection and returns the current element.
	//
	// Returns:
	//  - T: The current element in the collection.
	//  - error: An error if the iterator is exhausted or if an error occurred
	//    while consuming the element.
	Consume() (T, error)

	// Restart resets the iterator to the beginning of the
	// collection.
	Restart()
}
