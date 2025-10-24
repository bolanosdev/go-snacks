package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	filtered := list.Filter(func(it int) bool { return it == 1 })
	require.Equal(t, len(filtered), 1)
}

func TestFind(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}

	found, ok := list.Find(3)
	require.True(t, ok)
	require.Equal(t, 3, found)

	notFound, ok := list.Find(10)
	require.False(t, ok)
	require.Equal(t, 0, notFound)
}

func TestFindBy(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}

	found, ok := list.FindBy(func(it int) bool { return it == 3 })
	require.True(t, ok)
	require.Equal(t, 3, found)

	notFound, ok := list.FindBy(func(it int) bool { return it == 10 })
	require.False(t, ok)
	require.Equal(t, 0, notFound)
}

func TestFindIndex(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}

	index := list.FindIndex(func(it int) bool { return it == 3 })
	require.Equal(t, 2, index)

	index = list.FindIndex(func(it int) bool { return it == 10 })
	require.Equal(t, -1, index)
}

func TestSort(t *testing.T) {
	list := List[int]{5, 3, 1, 4, 2}
	sorted := list.Sort(func(a, b int) bool { return a < b })
	require.Equal(t, List[int]{1, 2, 3, 4, 5}, sorted)

	sorted = list.Sort(func(a, b int) bool { return a > b })
	require.Equal(t, List[int]{5, 4, 3, 2, 1}, sorted)
}

func TestAny(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	filtered, ok := list.Any(func(it int) bool { return it%2 == 0 })

	require.True(t, ok)
	require.Equal(t, len(filtered), 2)
}

func TestFirst(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	first, ok := list.First()
	require.True(t, ok)
	require.Equal(t, 1, first)

	empty_list := List[int]{}
	first, ok = empty_list.First()
	require.False(t, ok)
	require.Equal(t, 0, first)
}

func TestLast(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	last, ok := list.Last()
	require.True(t, ok)
	require.Equal(t, 5, last)

	empty_list := List[int]{}
	last, ok = empty_list.Last()
	require.False(t, ok)
	require.Equal(t, 0, last)
}

func TestReverse(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	reversed := list.Reverse()
	require.Equal(t, List[int]{5, 4, 3, 2, 1}, reversed)
}

func TestToMap(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	list := List[Person]{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	m := ToMap(list, func(p Person) int { return p.ID })

	require.Equal(t, 3, m.Len())
	val, ok := m.Get(1)
	require.True(t, ok)
	require.Equal(t, "Alice", val.Name)
}

func TestGroupBy(t *testing.T) {
	type Person struct {
		Age  int
		Name string
	}

	list := List[Person]{
		{Age: 25, Name: "Alice"},
		{Age: 30, Name: "Bob"},
		{Age: 25, Name: "Charlie"},
		{Age: 30, Name: "David"},
	}

	grouped := GroupBy(list, func(p Person) int { return p.Age })

	require.Equal(t, 2, grouped.Len())

	age25, ok := grouped.Get(25)
	require.True(t, ok)
	require.Equal(t, 2, len(age25))

	age30, ok := grouped.Get(30)
	require.True(t, ok)
	require.Equal(t, 2, len(age30))
}

func TestFold(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}

	sum := Fold(list, 0, func(acc int, item int) int {
		return acc + item
	})
	require.Equal(t, 15, sum)

	product := Fold(list, 1, func(acc int, item int) int {
		return acc * item
	})
	require.Equal(t, 120, product)

	str := Fold(List[string]{"a", "b", "c"}, "", func(acc string, item string) string {
		return acc + item
	})
	require.Equal(t, "abc", str)
}
