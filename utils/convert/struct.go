package convert

import (
	"fmt"
	"reflect"
)

// 把数组中的各个元素的指定filed取出来放到一个数组中
// 要求Array所有元素都是同一种类型，虽然支持指针混用但建议尽可能避免。
// ary := []interface{}{foo{ID: 1, Field: "foo1"}, &foo{ID: 2, Field: "foo22"}}
// fieldAry := Explode(ary,"Field") // []interface{}{"foo1","foo22"}
func ToExplodeStructField(array interface{}, field string) (explodeArray []interface{}, err error) {
	// 跳过Type的检查，异常自己处理=。=
	// 传入不是Array或者Slice的面壁去。。
	v := reflect.ValueOf(array)
	len := v.Len()
	explodeArray = make([]interface{}, len)
	var fieldIndex []int
	for i := 0; i < len; i++ {

		eV := v.Index(i).Elem()
		for eV.Kind() == reflect.Ptr {
			eV = eV.Elem()
		}

		// 缓存索引，加快取值速度
		if fieldIndex == nil {
			t := eV.Type()
			if fStruct, ok := t.FieldByName(field); ok {
				fieldIndex = fStruct.Index
			} else {
				return nil, fmt.Errorf("Invalid field : %s not exist in %v at array[%v]", field, eV, i)
			}
		}
		explodeArray[i] = eV.FieldByIndex(fieldIndex).Interface()
	}
	return
}

func MustExplodeStructField(array interface{}, field string) (explodeArray []interface{}) {
	if explodeArray, err := ToExplodeStructField(array, field); err != nil {
		panic(err)
	} else {
		return explodeArray
	}
}
