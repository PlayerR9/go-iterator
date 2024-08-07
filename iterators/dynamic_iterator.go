package iterators

// DynamicIterator is a struct that allows iterating over a collection
// of iterators of type Iterater[T].
type DynamicIterator[E, T any] struct {
	// source is the iterator over the collection of iterators.
	source Iterater[E]

	// iter is the iterator in the collection.
	iter Iterater[T]

	// transition is the transition function that takes an element of type E and
	// returns an iterator.
	transition func(E) Iterater[T]
}

// Consume implements the Iterater interface.
func (di *DynamicIterator[E, T]) Consume() (T, error) {
	if di.iter == nil {
		iter, err := di.source.Consume()
		if err != nil {
			return *new(T), err
		}

		di.iter = di.transition(iter)
	}

	var val T
	var err error

	for {
		val, err = di.iter.Consume()
		if err == nil {
			break
		}

		if err != Exhausted {
			return *new(T), err
		}

		iter, err := di.source.Consume()
		if err != nil {
			return *new(T), err
		}

		di.iter = di.transition(iter)
	}

	return val, nil
}

// Restart implements the Iterater interface.
func (di *DynamicIterator[E, T]) Restart() {
	di.iter = nil
	di.source.Restart()
}

// IteratorFromIterator creates a new iterator over a collection of iterators
// of type Iterater[T].
// It uses the input iterator to iterate over the collection of iterators and
// return the elements from each iterator in turn.
//
// Parameters:
//   - source: The iterator over the collection of iterators to iterate over.
//   - f: The transition function that takes an element of type E and returns
//     an iterator.
//
// Return:
//   - *DynamicIterator[E, T]: The new iterator. Nil if f or source is nil.
func NewDynamicIterator[E, T any](source Iterater[E], f func(E) Iterater[T]) *DynamicIterator[E, T] {
	if f == nil || source == nil {
		return nil
	}

	iter := &DynamicIterator[E, T]{
		source: source,
		iter:   nil,
	}

	iter.transition = f

	return iter
}
