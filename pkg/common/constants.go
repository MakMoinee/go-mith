package common

import (
	"reflect"
)

const (
	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json; charset=UTF-8"
)

var (
	SLICE_STRING []string
	SLICE_INT    []int
	INT          int
	STRING       string
)

var TypeValues = map[int]reflect.Type{
	1: reflect.ValueOf(SLICE_INT).Type(),
	2: reflect.ValueOf(SLICE_STRING).Type(),
	3: reflect.ValueOf(INT).Type(),
	4: reflect.ValueOf(STRING).Type(),
}

var ValueTypes = map[reflect.Type]reflect.Value{
	reflect.ValueOf(SLICE_INT).Type():    reflect.ValueOf(SLICE_INT),
	reflect.ValueOf(SLICE_STRING).Type(): reflect.ValueOf(SLICE_STRING),
	reflect.ValueOf(INT).Type():          reflect.ValueOf(INT),
	reflect.ValueOf(STRING).Type():       reflect.ValueOf(STRING),
}

func GetData(flag int) ([]int, []string, int, string) {
	sliceInt := []int{}
	sliceString := []string{}
	resultInt := 0
	resultString := ""

	if dataType, exist := TypeValues[flag]; exist {
		if values, existVal := ValueTypes[dataType]; existVal {
			switch flag {
			case 1:
				sliceInt = values.Interface().([]int)
			case 2:
				sliceString = values.Interface().([]string)
			case 3:
				resultInt = values.Interface().(int)
			case 4:
				resultString = values.Interface().(string)
			}
		}
	}

	return sliceInt, sliceString, resultInt, resultString
}

func InitializeMaps() {
	TypeValues = map[int]reflect.Type{
		1: reflect.ValueOf(SLICE_INT).Type(),
		2: reflect.ValueOf(SLICE_STRING).Type(),
		3: reflect.ValueOf(INT).Type(),
		4: reflect.ValueOf(STRING).Type(),
	}

	ValueTypes = map[reflect.Type]reflect.Value{
		reflect.ValueOf(SLICE_INT).Type():    reflect.ValueOf(SLICE_INT),
		reflect.ValueOf(SLICE_STRING).Type(): reflect.ValueOf(SLICE_STRING),
		reflect.ValueOf(INT).Type():          reflect.ValueOf(INT),
		reflect.ValueOf(STRING).Type():       reflect.ValueOf(STRING),
	}
}
