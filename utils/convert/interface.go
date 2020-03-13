package convert

import (
	"fmt"
	"reflect"
)

func MustInterfaceArray(array interface{}) []interface{} {
	if resArray, err := ToInterfaceArray(array); err == nil {
		return resArray
	} else {
		panic(err)
	}
}

func ToInterfaceArray(array interface{}) ([]interface{}, error) {
	t := reflect.TypeOf(array)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(array)
			resArray := make([]interface{}, v.Len())
			for index, _ := range resArray {
				resArray[index] = v.Index(index).Interface()
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", array)
}
