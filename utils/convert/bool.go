package convert

import "fmt"

func MustBool(value interface{}) bool {
	if res, err := ToBool(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 字符串中将"1", "t", "T", "true", "TRUE", "True"转为true
// 将"0", "f", "F", "false", "FALSE", "False", ""转为false，空字符串也转为false
// 数字中仅支持将0转为false，1转为true。
func ToBool(value interface{}) (res bool, err error) {
	res = false
	err = nil
	switch value.(type) {
	case string:
		{
			str := MustString(value)
			switch str {
			case "1", "t", "T", "true", "TRUE", "True":
				{
					res = true
				}
			case "0", "f", "F", "false", "FALSE", "False", "":
				{
					res = false
				}
			default:
				err = fmt.Errorf("convert: %v to boolean failed.", value)
			}
		}
	case bool:
		{
			res = value.(bool)
		}
	default:
		valueStr := MustString(value)
		res, err = ToBool(valueStr)
	}
	return
}