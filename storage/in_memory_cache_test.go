package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadAndGetMemoryCache(t *testing.T) {
	cache := NewInMemoryCache[string, string]()
	item, ok := cache.Get("foo")

	require.False(t, ok)
	require.Equal(t, item, "")

	cache.Set("foo", "bar")
	item, ok = cache.Get("foo")

	require.True(t, ok)
	require.Equal(t, item, "bar")

	item, ok = cache.Set("foo", "baz")
	require.True(t, ok)
	require.Equal(t, item, "baz")
}

func TestReadAndPopFromMemoryCache(t *testing.T) {
	cache := NewInMemoryCache[string, string]()
	item, ok := cache.Set("foo", "bar")
	require.True(t, ok)
	require.Equal(t, item, "bar")

	item, ok = cache.Pop("bar")
	require.False(t, ok)
	require.Equal(t, item, "")

	item, ok = cache.Pop("foo")
	require.True(t, ok)
	require.Equal(t, item, "bar")
}

func TestReadAndRemoveFromMemoryCache(t *testing.T) {
	cache := NewInMemoryCache[string, string]()
	item, ok := cache.Set("foo", "bar")
	require.True(t, ok)
	require.Equal(t, item, "bar")

	cache.Remove("bar")
	item, ok = cache.Get("foo")
	require.True(t, ok)
	require.Equal(t, item, "bar")

	cache.Remove("foo")
	item, ok = cache.Get("foo")
	require.False(t, ok)
	require.Equal(t, item, "")
}
