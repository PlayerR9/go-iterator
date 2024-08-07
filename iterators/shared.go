package iterators

import "errors"

// Exhausted is an error that is returned when an iterator is exhausted. Callers
// should return this error by itself and not wrap it as callers will test
// this error using ==.
//
// This error should only be used when signaling graceful termination.
var Exhausted error

func init() {
	Exhausted = errors.New("iterator is exhausted")
}

// Iterable is an interface that defines a method to get an iterator over a
// collection of elements of type T. It is implemented by data structures that
// can be iterated over.
type Iterable[T any] interface {
	// Iterator returns an iterator over the collection of elements.
	//
	// Returns:
	//   - Iterater[T]: An iterator over the collection of elements.
	Iterator() Iterater[T]
}

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
