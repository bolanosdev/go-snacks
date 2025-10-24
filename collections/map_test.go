package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapInit(t *testing.T) {
	m := Map[string, int]{}

	_, ok := m.Get("a")
	require.False(t, ok)
	require.Equal(t, 0, m.Len())
}

func TestMapGetSet(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2}

	val, ok := m.Get("a")
	require.True(t, ok)
	require.Equal(t, 1, val)

	m.Set("c", 3)
	val, ok = m.Get("c")
	require.True(t, ok)
	require.Equal(t, 3, val)
}

func TestMapRewrite(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2}

	val, ok := m.Get("a")
	require.True(t, ok)
	require.Equal(t, 1, val)

	m.Set("a", 3)
	val, ok = m.Get("a")
	require.True(t, ok)
	require.Equal(t, 3, val)
}

func TestMapHas(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2}

	require.True(t, m.Has("a"))
	require.False(t, m.Has("z"))
}

func TestMapDelete(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2}

	m.Delete("a")
	require.False(t, m.Has("a"))
	require.Equal(t, 1, m.Len())
}

func TestMapKeys(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2, "c": 3}

	keys := m.Keys()
	require.Equal(t, 3, len(keys))
	require.Contains(t, keys, "a")
	require.Contains(t, keys, "b")
	require.Contains(t, keys, "c")
}

func TestMapValues(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2, "c": 3}

	values := m.Values()
	require.Equal(t, 3, len(values))
	require.Contains(t, values, 1)
	require.Contains(t, values, 2)
	require.Contains(t, values, 3)
}

func TestMapClear(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2, "c": 3}

	m.Clear()
	require.Equal(t, 0, m.Len())
}

func TestMapCopy(t *testing.T) {
	m := Map[string, int]{"a": 1, "b": 2}

	copied := m.Copy()
	require.Equal(t, m.Len(), copied.Len())

	copied.Set("c", 3)
	require.False(t, m.Has("c"))
	require.True(t, copied.Has("c"))
}
