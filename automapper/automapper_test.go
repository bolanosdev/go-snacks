package automapper

import (
	"fmt"
	"testing"

	"github.com/bolanosdev/go-snacks/collections"
	"github.com/stretchr/testify/require"
)

func TestAutoMapperMap(t *testing.T) {
	mapper := New()
	err := mapper.AddMapper(func(it int) (*string, error) {
		result := fmt.Sprintf("value-%d", it)
		return &result, nil
	})
	require.NoError(t, err)

	var result string
	err = mapper.Map(5, &result)
	require.NoError(t, err)
	require.Equal(t, "value-5", result)
}

func TestAutoMapperAddError(t *testing.T) {
	mapper := New()
	err := mapper.AddMapper(func(it int) (*string, error) {
		result := "value"
		return &result, nil
	})
	require.NoError(t, err)

	err = mapper.AddMapper(func(it int) (*string, error) {
		result := "other"
		return &result, nil
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "mapper already registered")
}

func TestAutoMapperMapList(t *testing.T) {
	mapper := New()
	err := mapper.AddMapper(func(it int) (*string, error) {
		result := fmt.Sprintf("value-%d", it)
		return &result, nil
	})
	require.NoError(t, err)

	source := []int{1, 2, 3, 4, 5}
	var result []string
	err = mapper.MapList(source, &result)
	require.NoError(t, err)
	require.Equal(t, []string{"value-1", "value-2", "value-3", "value-4", "value-5"}, result)
}

func TestAutoMapperMapListWithValues(t *testing.T) {
	mapper := New()
	err := mapper.AddMapper(func(it int) (*string, error) {
		result := fmt.Sprintf("value-%d", it)
		return &result, nil
	})
	require.NoError(t, err)

	list := collections.List[int]{1, 2, 3, 4, 5}
	var result []string
	err = mapper.MapList(list.Values(), &result)
	require.NoError(t, err)
	require.Equal(t, []string{"value-1", "value-2", "value-3", "value-4", "value-5"}, result)
}

func TestAutoMapperConfigure(t *testing.T) {
	configure := func(m *AutoMapper) error {
		return m.AddMapper(func(it int) (*string, error) {
			result := fmt.Sprintf("value-%d", it)
			return &result, nil
		})
	}

	mapper, err := New().Configure(configure)
	require.NoError(t, err)

	var result string
	err = mapper.Map(5, &result)
	require.NoError(t, err)
	require.Equal(t, "value-5", result)
}

func TestAutoMapperConfigureError(t *testing.T) {
	configure := func(m *AutoMapper) error {
		if err := m.AddMapper(func(it int) (*string, error) {
			result := "value"
			return &result, nil
		}); err != nil {
			return err
		}
		return m.AddMapper(func(it int) (*string, error) {
			result := "other"
			return &result, nil
		})
	}

	_, err := New().Configure(configure)
	require.Error(t, err)
	require.Contains(t, err.Error(), "mapper already registered")
}

func TestAutoMapperMapperError(t *testing.T) {
	mapper := New()
	err := mapper.AddMapper(func(it int) (*string, error) {
		if it < 0 {
			return nil, fmt.Errorf("negative value not allowed: %d", it)
		}
		result := fmt.Sprintf("value-%d", it)
		return &result, nil
	})
	require.NoError(t, err)

	var result string
	err = mapper.Map(-5, &result)
	require.Error(t, err)
	require.Contains(t, err.Error(), "negative value not allowed")

	err = mapper.Map(5, &result)
	require.NoError(t, err)
	require.Equal(t, "value-5", result)
}

func TestAutoMapperMapListWithError(t *testing.T) {
	mapper := New()
	err := mapper.AddMapper(func(it int) (*string, error) {
		if it < 0 {
			return nil, fmt.Errorf("negative value not allowed: %d", it)
		}
		result := fmt.Sprintf("value-%d", it)
		return &result, nil
	})
	require.NoError(t, err)

	source := []int{1, 2, -3, 4, 5}
	var result []string
	err = mapper.MapList(source, &result)
	require.Error(t, err)
	require.Contains(t, err.Error(), "negative value not allowed")
}

func TestAutoMapperMapWithNilReturn(t *testing.T) {
	mapper := New()
	err := mapper.AddMapper(func(it int) (*string, error) {
		if it == 0 {
			return nil, nil
		}
		result := fmt.Sprintf("value-%d", it)
		return &result, nil
	})
	require.NoError(t, err)

	var result string
	err = mapper.Map(0, &result)
	require.NoError(t, err)
	require.Equal(t, "", result)

	err = mapper.Map(5, &result)
	require.NoError(t, err)
	require.Equal(t, "value-5", result)
}
