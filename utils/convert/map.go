package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 把map转成map[string]interface{}，key的值使用MustString计算。
// 如果子项中也有map，则继续递归执行直到全部转换为map[string]interface{}
// 如果子项有[]interface{}，则要继续判定slice的元素中的类型
// 常用于各种xml\yaml\json转换为map的结果的统一处理。
func MustMapStringInterfaceRecursions(leafMap interface{}) map[string]interface{} {
	leafType := reflect.TypeOf(leafMap)
	if leafType.Kind() != reflect.Map {
		return nil
	}
	leafValue := reflect.ValueOf(leafMap)
	if leafValue.Len() == 0 {
		return nil
	}

	resMap := make(map[string]interface{})
	leafKeyValues := leafValue.MapKeys()
	// key的value
	for _, leafKeyValue := range leafKeyValues {
		// node的value
		nodeValue := leafValue.MapIndex(leafKeyValue)

		// 获得实际的key和node
		k := leafKeyValue.Interface()
		node := nodeValue.Interface()
		if nodeValue.IsNil() {
			continue
		}

		strKey := MustString(k)
		nodeType := reflect.TypeOf(node)

		switch nodeType.Kind() {
		case reflect.Map:
			temp := MustMapStringInterfaceRecursions(node)
			if temp != nil {
				resMap[strKey] = temp
			}
		case reflect.Slice, reflect.Array:
			temp := MustMapStringInterfaceRecursionsInArrayInterface(node)
			if temp != nil {
				resMap[strKey] = temp
			}
		default:
			resMap[strKey] = node
		}
	}
	return resMap
}

// 协助处理[]interface{}中的map[interface{}]interface{}为map[string]interface{}
func MustMapStringInterfaceRecursionsInArrayInterface(leafAry interface{}) []interface{} {
	leafType := reflect.TypeOf(leafAry)
	if leafType.Kind() != reflect.Array &&
		leafType.Kind() != reflect.Slice {
		return nil
	}
	leafValue := reflect.ValueOf(leafAry)
	if leafValue.Len() == 0 {
		return nil
	}
	resAry := make([]interface{}, 0)

	for i := 0; i < leafValue.Len(); i++ {
		nodeValue := leafValue.Index(i)
		// 获得实际的key和node

		node := nodeValue.Interface()
		nodeType := reflect.TypeOf(node)

		switch nodeType.Kind() {
		case reflect.Array, reflect.Slice:
			temp := MustMapStringInterfaceRecursionsInArrayInterface(node)
			if temp != nil {
				resAry = append(resAry, temp)
			}
		case reflect.Map:
			temp := MustMapStringInterfaceRecursions(node)
			if temp != nil {
				resAry = append(resAry, temp)
			}
		default:
			resAry = append(resAry, node)
		}
	}

	return resAry
}

func ToArrayMapString(maps []map[string]interface{}) []map[string]string {
	records := []map[string]string{}
	for i := 0; i < len(maps); i++ {
		m := ToMapString(maps[i])
		records = append(records, m)
	}
	return records
}

func ToArrayMapStr(m []map[string]interface{}) (string, error) {
	if m == nil {
		return "", nil
	}
	bt, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bt[:]), nil
}

func ToMapStr(m map[string]interface{}) (string, error) {
	if m == nil {
		return "", nil
	}
	bt, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bt[:]), nil
}

func ToMapByte(m map[string]interface{}) []byte {
	if m == nil {
		return nil
	}
	bt, err := json.Marshal(m)
	if err != nil {
		println(err.Error())
		return nil
	}
	return bt[:]
}

func ToObjStr(o interface{}) (string) {
	if o == nil {
		return ""
	}
	bt, err := json.Marshal(o)
	if err != nil {
		println(err.Error())
		return ""
	}
	return string(bt[:])
}

func ToInterfaceByte(m interface{}) []byte {
	if m == nil {
		return nil
	}
	bt, err := json.Marshal(m)
	if err != nil {
		println(err.Error())
		return nil
	}
	return bt[:]
}

func ToMapString(m map[string]interface{}) map[string]string {
	record := make(map[string]string)
	for k, v := range m {
		val := fmt.Sprintf("%v", v)
		record[k] = val
	}
	return record

}
func Obj2Map(obj interface{}) (map[string]interface{}, error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		panic(err)
		return nil, err
	}
	return result, nil
}

func Obj2MapString(obj interface{}) (map[string]string, error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	if err := json.Unmarshal(b, &result); err != nil {
		panic(err)
		return nil, err
	}
	return result, nil
}

func Obj2ListMapString(obj interface{}) ([]map[string]string, error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var result []map[string]string
	if err := json.Unmarshal(b, &result); err != nil {
		panic(err)
		return nil, err
	}
	return result, nil
}

func Obj2ListMap(obj interface{}) ([]map[string]interface{}, error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		panic(err)
		return nil, err
	}
	return result, nil
}

func Byte2Map(b []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func String2Map(str string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func String2MapList(str string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}
