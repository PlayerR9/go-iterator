package slice

import "io"

// SliceIterator is a struct that allows iterating over a collection of
// iterators of type Iterater[T].
type SliceIterator[T any] struct {
	// source is the iterator over the collection of iterators.
	source Iterater[[]T]

	// iter is the iterator in the collection.
	iter *SimpleIterator[T]
}

// Consume implements the Iterater interface.
func (pi *SliceIterator[T]) Consume() (T, error) {
	if pi.iter == nil {
		values, err := pi.source.Consume()
		if err != nil {
			return *new(T), err
		}

		pi.iter = NewSimpleIterator(values)
	}

	var val T
	var err error

	for {
		val, err = pi.iter.Consume()
		if err == nil {
			break
		}

		if err != io.EOF {
			return *new(T), err
		}

		iter, err := pi.source.Consume()
		if err != nil {
			return *new(T), err
		}

		pi.iter = NewSimpleIterator(iter)
	}

	return val, nil
}

// Restart implements the Iterater interface.
func (pi *SliceIterator[T]) Restart() {
	pi.iter = nil
	pi.source.Restart()
}

// IteratorFromIterator creates a new iterator over a collection of iterators
// of type Iterater[T].
// It uses the input iterator to iterate over the collection of iterators and
// return the elements from each iterator in turn.
//
// Parameters:
//   - source: The iterator over the collection of iterators to iterate over.
//
// Return:
//   - *SliceIterator[T]: The new iterator over the collection of elements.
//     Nil if source is nil.
func NewSliceIterator[T any](source Iterater[[]T]) *SliceIterator[T] {
	if source == nil {
		return nil
	}

	pi := &SliceIterator[T]{
		source: source,
		iter:   nil,
	}

	return pi
}
