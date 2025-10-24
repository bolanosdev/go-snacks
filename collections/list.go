package collections

import "sort"

type List[T any] []T

func (l List[T]) Len() int {
	return len(l)
}

func (l List[T]) Get(index int) T {
	return l[index]
}

func (s List[T]) Filter(predicate func(T) bool) List[T] {
	result := make(List[T], 0)
	for _, item := range s {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Find returns the first element that equals the given value and true,
// or the zero value and false if not found
func (s List[T]) Find(value T) (T, bool) {
	return s.Filter(func(item T) bool {
		return any(item) == any(value)
	}).First()
}

// Find returns the first element that satisfies the predicate and true,
// or the zero value and false if not found
func (s List[T]) FindBy(predicate func(T) bool) (T, bool) {
	return s.Filter(predicate).First()
}

// FindIndex returns the index of the first element that satisfies the predicate,
// or -1 if not found
func (s List[T]) FindIndex(predicate func(T) bool) int {
	for i, item := range s {
		if predicate(item) {
			return i
		}
	}
	return -1
}

func (s List[T]) Sort(less func(T, T) bool) List[T] {
	result := make(List[T], len(s))
	copy(result, s)

	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// Any returns the first element that satisfies the predicate and true,
// or nil and false if not found
func (s List[T]) Any(predicate func(T) bool) (List[T], bool) {
	filtered := s.Filter(predicate)
	if len(filtered) > 0 {
		return filtered, true
	}
	return nil, false
}

// First returns the first element and true, or zero value and false if empty
func (s List[T]) First() (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}

// Last returns the last element and true, or zero value and false if empty
func (s List[T]) Last() (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[len(s)-1], true
}

// Reverse returns a new slice with elements in reverse order
func (s List[T]) Reverse() List[T] {
	result := make(List[T], len(s))
	for i, item := range s {
		result[len(s)-1-i] = item
	}
	return result
}

func (s List[T]) Values() []T {
	return []T(s)
}

func ToMap[T any, K comparable](list List[T], keyFunc func(T) K) Map[K, T] {
	result := make(Map[K, T])
	for _, item := range list {
		result[keyFunc(item)] = item
	}
	return result
}

func GroupBy[T any, K comparable](list List[T], keyFunc func(T) K) Map[K, List[T]] {
	result := make(Map[K, List[T]])
	for _, item := range list {
		key := keyFunc(item)
		result[key] = append(result[key], item)
	}
	return result
}

func Fold[T any, R any](list List[T], initial R, fn func(R, T) R) R {
	acc := initial
	for _, item := range list {
		acc = fn(acc, item)
	}
	return acc
}
