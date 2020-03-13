package convert

import (
	"fmt"
	"reflect"
	"strings"
)

// 强制转换一个对象为string
func MustString(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

// 转换一个对象为string
func ToString(value interface{}) string {
	// fake方法用于统一接口
	return MustString(value)
}

// 强转一个字符串数组
func MustStringArray(array interface{}) (resArray []string) {
	t := reflect.TypeOf(array)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(array)
			resArray = make([]string, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustString(v.Index(index).Interface())
			}
			return
		}
	}
	return []string{MustString(array)}
}

// 尝试一个字符串数组
func ToStringArray(value interface{}) ([]string, error) {
	return MustStringArray(value), nil
}

// 尝试一个字符串数组
func ToStringArrayInt64(value, sep string) (arr []int64) {
	if value == "" {
		return
	}
	str := strings.Split(value, sep)
	for _, v := range str {
		arr = append(arr, MustInt64(v))
	}
	return
}
