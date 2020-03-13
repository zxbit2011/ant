package convert


import (
	"fmt"
	"reflect"
)

// 尽最大努力将目标转换为[]int
func ToIntArray(value interface{}) (resArray []int, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToInt(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

// 尽最大努力将目标转换为[]int8
func ToInt8Array(value interface{}) (resArray []int8, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int8, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToInt8(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

// 尽最大努力将目标转换为[]int16
func ToInt16Array(value interface{}) (resArray []int16, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int16, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToInt16(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

// 尽最大努力将目标转换为[]int32
func ToInt32Array(value interface{}) (resArray []int32, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int32, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToInt32(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)

}

// 尽最大努力将目标转换为[]int64
func ToInt64Array(value interface{}) (resArray []int64, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int64, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToInt64(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)

}

//尽最大努力将目标转换为[]uint
func ToUintArray(value interface{}) (resArray []uint, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToUint(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

//尽最大努力将目标转换为[]uint8
func ToUint8Array(value interface{}) (resArray []uint8, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint8, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToUint8(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

//尽最大努力将目标转换为[]uint16
func ToUint16Array(value interface{}) (resArray []uint16, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint16, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToUint16(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

//尽最大努力将目标转换为[]uint32
func ToUint32Array(value interface{}) (resArray []uint32, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint32, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToUint32(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

//尽最大努力将目标转换为[]uint64
func ToUint64Array(value interface{}) (resArray []uint64, err error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint64, v.Len())
			for index, _ := range resArray {
				resArray[index], err = ToUint64(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice.", value)
}

// 尽最大努力将目标转换为[]int，失败会panic
func MustIntArray(value interface{}) (resArray []int) {
	if res, err := ToIntArray(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 尽最大努力将目标转换为[]int8，失败会panic
func MustInt8Array(value interface{}) (resArray []int8) {
	if res, err := ToInt8Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 尽最大努力将目标转换为[]int16，失败会panic
func MustInt16Array(value interface{}) (resArray []int16) {
	if res, err := ToInt16Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 尽最大努力将目标转换为[]int32，失败会panic
func MustInt32Array(value interface{}) (resArray []int32) {
	if res, err := ToInt32Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 尽最大努力将目标转换为[]int64，失败会panic
func MustInt64Array(value interface{}) (resArray []int64) {
	if res, err := ToInt64Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

//尽最大努力将目标转换为[]uint,失败会panic
func MustUintArray(value interface{}) (resArray []uint) {
	if res, err := ToUintArray(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

//尽最大努力将目标转换为[]uint8,失败会panic
func MustUint8Array(value interface{}) (resArray []uint8) {
	if res, err := ToUint8Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

//尽最大努力将目标转换为[]uint16,失败会panic
func MustUint16Array(value interface{}) (resArray []uint16) {
	if res, err := ToUint16Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

//尽最大努力将目标转换为[]uint32,失败会panic
func MustUint32Array(value interface{}) (resArray []uint32) {
	if res, err := ToUint32Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

//尽最大努力将目标转换为[]uint64,失败会panic
func MustUint64Array(value interface{}) (resArray []uint64) {
	if res, err := ToUint64Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}
