package main

type FilteredSlice[T any] struct {
	slice     []T
	predicate func(T) bool
}

func NewFilteredSlice[T any](slice []T, predicate func(T) bool) *FilteredSlice[T] {
	return &FilteredSlice[T]{
		slice:     slice,
		predicate: predicate,
	}
}

func (fs *FilteredSlice[T]) Keep() []T {
	result := make([]T, 0, len(fs.slice))
	for _, v := range fs.slice {
		if fs.predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func (fs *FilteredSlice[T]) Discard() []T {
	result := make([]T, 0, len(fs.slice))
	for _, v := range fs.slice {
		if !fs.predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {

}
