package automapper

import (
	"errors"
	"reflect"
)

type mapperFunc struct {
	fn interface{}
}

type AutoMapper struct {
	mappers map[string]mapperFunc
}

func New() *AutoMapper {
	return &AutoMapper{
		mappers: make(map[string]mapperFunc),
	}
}

func (m *AutoMapper) Configure(configure func(*AutoMapper) error) (*AutoMapper, error) {
	if err := configure(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *AutoMapper) AddMapper(mapper interface{}) error {
	mapperType := reflect.TypeOf(mapper)
	if mapperType.Kind() != reflect.Func {
		return errors.New("mapper must be a function")
	}

	if mapperType.NumIn() != 1 || mapperType.NumOut() != 1 {
		return errors.New("mapper must have exactly one input and one output")
	}

	sourceType := mapperType.In(0)
	destType := mapperType.Out(0)
	key := sourceType.String() + "->" + destType.String()

	if _, exists := m.mappers[key]; exists {
		return errors.New("mapper already registered for this type pair: " + key)
	}

	m.mappers[key] = mapperFunc{fn: mapper}
	return nil
}

func (m *AutoMapper) Map(source interface{}, dest interface{}) error {
	sourceType := reflect.TypeOf(source)
	destType := reflect.TypeOf(dest)

	if destType.Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer")
	}

	destElemType := destType.Elem()
	key := sourceType.String() + "->" + destElemType.String()

	mapperFunc, exists := m.mappers[key]
	if !exists {
		return errors.New("no mapper found for type pair: " + key)
	}

	fn := reflect.ValueOf(mapperFunc.fn)
	result := fn.Call([]reflect.Value{reflect.ValueOf(source)})

	reflect.ValueOf(dest).Elem().Set(result[0])
	return nil
}

func (m *AutoMapper) MapList(source interface{}, dest interface{}) error {
	sourceVal := reflect.ValueOf(source)
	destVal := reflect.ValueOf(dest)

	if sourceVal.Kind() != reflect.Slice {
		return errors.New("source must be a slice")
	}

	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Slice {
		return errors.New("destination must be a pointer to a slice")
	}

	destSlice := destVal.Elem()
	destSliceType := destSlice.Type().Elem()

	result := reflect.MakeSlice(reflect.SliceOf(destSliceType), sourceVal.Len(), sourceVal.Len())

	for i := 0; i < sourceVal.Len(); i++ {
		sourceItem := sourceVal.Index(i).Interface()
		destItem := reflect.New(destSliceType)

		if err := m.Map(sourceItem, destItem.Interface()); err != nil {
			return err
		}

		result.Index(i).Set(destItem.Elem())
	}

	destSlice.Set(result)
	return nil
}
