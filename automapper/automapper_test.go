package automapper

import (
	"fmt"
	"testing"

	"github.com/bolanosdev/go-snacks/collections"
	"github.com/stretchr/testify/require"
)

func TestAutoMapperMap(t *testing.T) {
	mapper := NewAutoMapper()
	err := mapper.AddMapper(func(it int) string { return fmt.Sprintf("value-%d", it) })
	require.NoError(t, err)

	var result string
	err = mapper.Map(5, &result)
	require.NoError(t, err)
	require.Equal(t, "value-5", result)
}

func TestAutoMapperAddError(t *testing.T) {
	mapper := NewAutoMapper()
	err := mapper.AddMapper(func(it int) string { return "value" })
	require.NoError(t, err)

	err = mapper.AddMapper(func(it int) string { return "other" })
	require.Error(t, err)
	require.Contains(t, err.Error(), "mapper already registered")
}

func TestAutoMapperMapList(t *testing.T) {
	mapper := NewAutoMapper()
	err := mapper.AddMapper(func(it int) string { return fmt.Sprintf("value-%d", it) })
	require.NoError(t, err)

	source := []int{1, 2, 3, 4, 5}
	var result []string
	err = mapper.MapList(source, &result)
	require.NoError(t, err)
	require.Equal(t, []string{"value-1", "value-2", "value-3", "value-4", "value-5"}, result)
}

func TestAutoMapperMapListWithValues(t *testing.T) {
	mapper := NewAutoMapper()
	err := mapper.AddMapper(func(it int) string { return fmt.Sprintf("value-%d", it) })
	require.NoError(t, err)

	list := collections.List[int]{1, 2, 3, 4, 5}
	var result []string
	err = mapper.MapList(list.Values(), &result)
	require.NoError(t, err)
	require.Equal(t, []string{"value-1", "value-2", "value-3", "value-4", "value-5"}, result)
}
