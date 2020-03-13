package convert

import (
	"fmt"
	"strconv"
)

// 尽最大努力将一个值转为int类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt(value interface{}) (res int, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to int failed", value.(string))
						} else {
							res, err = ToInt(resBool)
						}
					} else {
						res, err = ToInt(resF64)
					}
				} else {
					res, err = ToInt(resU64)
				}

			} else {
				res = int(res64)
			}
		}
	case int:
		{
			res = int(value.(int))
		}
	case int8:
		{
			res = int(value.(int8))
		}
	case int16:
		{
			res = int(value.(int16))
		}
	case int32:
		{
			res = int(value.(int32))
		}
	case int64:
		{
			res = int(value.(int64))
		}
	case uint:
		{
			res = int(value.(uint))
		}
	case uint8:
		{
			res = int(value.(uint8))
		}
	case uint16:
		{
			res = int(value.(uint16))
		}
	case uint32:
		{
			res = int(value.(uint32))
		}
	case uint64:
		{
			res = int(value.(uint64))
		}
	case uintptr:
		{
			res = int(value.(uintptr))
		}

	case float32:
		{
			res = int(value.(float32))
		}
	case float64:
		{
			res = int(value.(float64))
		}

	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为int8类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt8(value interface{}) (res int8, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to int8 failed", value.(string))
						} else {
							res, err = ToInt8(resBool)
						}
					} else {
						res, err = ToInt8(resF64)
					}
				} else {
					res, err = ToInt8(resU64)
				}

			} else {
				res = int8(res64)
			}
		}
	case int:
		{
			res = int8(value.(int))
		}
	case int8:
		{
			res = int8(value.(int8))
		}
	case int16:
		{
			res = int8(value.(int16))
		}
	case int32:
		{
			res = int8(value.(int32))
		}
	case int64:
		{
			res = int8(value.(int64))
		}
	case uint:
		{
			res = int8(value.(uint))
		}
	case uint8:
		{
			res = int8(value.(uint8))
		}
	case uint16:
		{
			res = int8(value.(uint16))
		}
	case uint32:
		{
			res = int8(value.(uint32))
		}
	case uint64:
		{
			res = int8(value.(uint64))
		}
	case uintptr:
		{
			res = int8(value.(uintptr))
		}

	case float32:
		{
			res = int8(value.(float32))
		}
	case float64:
		{
			res = int8(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt8(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为int16类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt16(value interface{}) (res int16, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to int16 failed", value.(string))
						} else {
							res, err = ToInt16(resBool)
						}
					} else {
						res, err = ToInt16(resF64)
					}
				} else {
					res, err = ToInt16(resU64)
				}

			} else {
				res = int16(res64)
			}
		}
	case int:
		{
			res = int16(value.(int))
		}
	case int8:
		{
			res = int16(value.(int8))
		}
	case int16:
		{
			res = int16(value.(int16))
		}
	case int32:
		{
			res = int16(value.(int32))
		}
	case int64:
		{
			res = int16(value.(int64))
		}
	case uint:
		{
			res = int16(value.(uint))
		}
	case uint8:
		{
			res = int16(value.(uint8))
		}
	case uint16:
		{
			res = int16(value.(uint16))
		}
	case uint32:
		{
			res = int16(value.(uint32))
		}
	case uint64:
		{
			res = int16(value.(uint64))
		}
	case uintptr:
		{
			res = int16(value.(uintptr))
		}

	case float32:
		{
			res = int16(value.(float32))
		}
	case float64:
		{
			res = int16(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{

			valueStr := MustString(value)
			res, err = ToInt16(valueStr)

		}
	}
	return
}

// 尽最大努力将一个值转为int32类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt32(value interface{}) (res int32, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to int32 failed", value.(string))
						} else {
							res, err = ToInt32(resBool)
						}
					} else {
						res, err = ToInt32(resF64)
					}
				} else {
					res, err = ToInt32(resU64)
				}

			} else {
				res = int32(res64)
			}
		}
	case int:
		{
			res = int32(value.(int))
		}
	case int8:
		{
			res = int32(value.(int8))
		}
	case int16:
		{
			res = int32(value.(int16))
		}
	case int32:
		{
			res = int32(value.(int32))
		}
	case int64:
		{
			res = int32(value.(int64))
		}
	case uint:
		{
			res = int32(value.(uint))
		}
	case uint8:
		{
			res = int32(value.(uint8))
		}
	case uint16:
		{
			res = int32(value.(uint16))
		}
	case uint32:
		{
			res = int32(value.(uint32))
		}
	case uint64:
		{
			res = int32(value.(uint64))
		}
	case uintptr:
		{
			res = int32(value.(uintptr))
		}

	case float32:
		{
			res = int32(value.(float32))
		}
	case float64:
		{
			res = int32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt32(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为int64类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt64(value interface{}) (res int64, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to int64 failed", value.(string))
						} else {
							res, err = ToInt64(resBool)
						}
					} else {
						res, err = ToInt64(resF64)
					}
				} else {
					res, err = ToInt64(resU64)
				}

			} else {
				res = res64
			}
		}
	case int:
		{
			res = int64(value.(int))
		}
	case int8:
		{
			res = int64(value.(int8))
		}
	case int16:
		{
			res = int64(value.(int16))
		}
	case int32:
		{
			res = int64(value.(int64))
		}
	case uint:
		{
			res = int64(value.(uint))
		}
	case uint8:
		{
			res = int64(value.(uint8))
		}
	case uint16:
		{
			res = int64(value.(uint16))
		}
	case uint32:
		{
			res = int64(value.(uint32))
		}
	case uint64:
		{
			res = int64(value.(uint64))
		}
	case uintptr:
		{
			res = int64(value.(uintptr))
		}

	case float32:
		{
			res = int64(value.(float32))
		}
	case float64:
		{
			res = int64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt64(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为uint类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint(value interface{}) (res uint, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to Uint failed", value.(string))
						} else {
							res, err = ToUint(resBool)
						}
					} else {
						res, err = ToUint(resF64)
					}
				} else {
					res, err = ToUint(res64)
				}
			} else {
				res = uint(resU64)
			}
		}
	case int:
		{
			res = uint(value.(int))
		}
	case int8:
		{
			res = uint(value.(int8))
		}
	case int16:
		{
			res = uint(value.(int16))
		}
	case int32:
		{
			res = uint(value.(int32))
		}
	case int64:
		{
			res = uint(value.(int64))
		}

	case uint:
		{
			res = uint(value.(uint))
		}
	case uint8:
		{
			res = uint(value.(uint8))
		}
	case uint16:
		{
			res = uint(value.(uint16))
		}
	case uint32:
		{
			res = uint(value.(uint32))
		}
	case uint64:
		{
			res = uint(value.(uint64))
		}
	case uintptr:
		{
			res = uint(value.(uintptr))
		}

	case float32:
		{
			res = uint(value.(float32))
		}
	case float64:
		{
			res = uint(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为uint8类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint8(value interface{}) (res uint8, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to Uint8 failed", value.(string))
						} else {
							res, err = ToUint8(resBool)
						}
					} else {
						res, err = ToUint8(resF64)
					}
				} else {
					res, err = ToUint8(res64)
				}
			} else {
				res = uint8(resU64)
			}
		}
	case int:
		{
			res = uint8(value.(int))
		}
	case int8:
		{
			res = uint8(value.(int8))
		}
	case int16:
		{
			res = uint8(value.(int16))
		}
	case int32:
		{
			res = uint8(value.(int32))
		}
	case int64:
		{
			res = uint8(value.(int64))
		}

	case uint:
		{
			res = uint8(value.(uint))
		}
	case uint8:
		{
			res = uint8(value.(uint8))
		}
	case uint16:
		{
			res = uint8(value.(uint16))
		}
	case uint32:
		{
			res = uint8(value.(uint32))
		}
	case uint64:
		{
			res = uint8(value.(uint64))
		}
	case uintptr:
		{
			res = uint8(value.(uintptr))
		}

	case float32:
		{
			res = uint8(value.(float32))
		}
	case float64:
		{
			res = uint8(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint8(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为uint16类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint16(value interface{}) (res uint16, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to Uint16 failed", value.(string))
						} else {
							res, err = ToUint16(resBool)
						}
					} else {
						res, err = ToUint16(resF64)
					}
				} else {
					res, err = ToUint16(res64)
				}
			} else {
				res = uint16(resU64)
			}
		}
	case int:
		{
			res = uint16(value.(int))
		}
	case int8:
		{
			res = uint16(value.(int8))
		}
	case int16:
		{
			res = uint16(value.(int16))
		}
	case int32:
		{
			res = uint16(value.(int32))
		}
	case int64:
		{
			res = uint16(value.(int64))
		}

	case uint:
		{
			res = uint16(value.(uint))
		}
	case uint8:
		{
			res = uint16(value.(uint8))
		}
	case uint16:
		{
			res = uint16(value.(uint16))
		}
	case uint32:
		{
			res = uint16(value.(uint32))
		}
	case uint64:
		{
			res = uint16(value.(uint64))
		}
	case uintptr:
		{
			res = uint16(value.(uintptr))
		}

	case float32:
		{
			res = uint16(value.(float32))
		}
	case float64:
		{
			res = uint16(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint16(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为uint32类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint32(value interface{}) (res uint32, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to Uint32 failed", value.(string))
						} else {
							res, err = ToUint32(resBool)
						}
					} else {
						res, err = ToUint32(resF64)
					}
				} else {
					res, err = ToUint32(res64)
				}
			} else {
				res = uint32(resU64)
			}
		}
	case int:
		{
			res = uint32(value.(int))
		}
	case int8:
		{
			res = uint32(value.(int8))
		}
	case int16:
		{
			res = uint32(value.(int16))
		}
	case int32:
		{
			res = uint32(value.(int32))
		}
	case int64:
		{
			res = uint32(value.(int64))
		}

	case uint:
		{
			res = uint32(value.(uint))
		}
	case uint8:
		{
			res = uint32(value.(uint8))
		}
	case uint16:
		{
			res = uint32(value.(uint16))
		}
	case uint32:
		{
			res = uint32(value.(uint32))
		}
	case uint64:
		{
			res = uint32(value.(uint64))
		}
	case uintptr:
		{
			res = uint32(value.(uintptr))
		}

	case float32:
		{
			res = uint32(value.(float32))
		}
	case float64:
		{
			res = uint32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint32(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为uint64类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint64(value interface{}) (res uint64, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to Uint64 failed", value.(string))
						} else {
							res, err = ToUint64(resBool)
						}
					} else {
						res, err = ToUint64(resF64)
					}
				} else {
					res, err = ToUint64(res64)
				}
			} else {
				res = resU64
			}
		}
	case int:
		{
			res = uint64(value.(int))
		}
	case int8:
		{
			res = uint64(value.(int8))
		}
	case int16:
		{
			res = uint64(value.(int16))
		}
	case int32:
		{
			res = uint64(value.(int32))
		}
	case int64:
		{
			res = uint64(value.(int64))
		}

	case uint:
		{
			res = uint64(value.(uint))
		}
	case uint8:
		{
			res = uint64(value.(uint8))
		}
	case uint16:
		{
			res = uint64(value.(uint16))
		}
	case uint32:
		{
			res = uint64(value.(uint32))
		}
	case uint64:
		{
			res = uint64(value.(uint64))
		}
	case uintptr:
		{
			res = uint64(value.(uintptr))
		}

	case float32:
		{
			res = uint64(value.(float32))
		}
	case float64:
		{
			res = uint64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint64(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为float32类型的数据
// string会按顺序尝试将数据解析为float64\uint64\bool，然后再转换为float64
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToFloat32(value interface{}) (res float32, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {

						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to Uint64 failed", value.(string))
						} else {
							res, err = ToFloat32(resBool)
						}
					} else {
						res, err = ToFloat32(res64)
					}
				} else {
					res, err = ToFloat32(resU64)
				}
			} else {
				res, err = ToFloat32(resF64)
			}
		}
	case int:
		{
			res = float32(value.(int))
		}
	case int8:
		{
			res = float32(value.(int8))
		}
	case int16:
		{
			res = float32(value.(int16))
		}
	case int32:
		{
			res = float32(value.(int32))
		}
	case int64:
		{
			res = float32(value.(int64))
		}

	case uint:
		{
			res = float32(value.(uint))
		}
	case uint8:
		{
			res = float32(value.(uint8))
		}
	case uint16:
		{
			res = float32(value.(uint16))
		}
	case uint32:
		{
			res = float32(value.(uint32))
		}
	case uint64:
		{
			res = float32(value.(uint64))
		}
	case uintptr:
		{
			res = float32(value.(uintptr))
		}

	case float32:
		{
			res = float32(value.(float32))
		}
	case float64:
		{
			res = float32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToFloat32(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为float64类型的数据
// string会按顺序尝试将数据解析为float64\uint64\bool，然后再转换为float64
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToFloat64(value interface{}) (res float64, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {

						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("convert: string \"%s\" to Uint64 failed", value.(string))
						} else {
							res, err = ToFloat64(resBool)
						}
					} else {
						res, err = ToFloat64(res64)
					}
				} else {
					res, err = ToFloat64(resU64)
				}
			} else {
				res = resF64
			}
		}
	case int:
		{
			res = float64(value.(int))
		}
	case int8:
		{
			res = float64(value.(int8))
		}
	case int16:
		{
			res = float64(value.(int16))
		}
	case int32:
		{
			res = float64(value.(int32))
		}
	case int64:
		{
			res = float64(value.(int64))
		}

	case uint:
		{
			res = float64(value.(uint))
		}
	case uint8:
		{
			res = float64(value.(uint8))
		}
	case uint16:
		{
			res = float64(value.(uint16))
		}
	case uint32:
		{
			res = float64(value.(uint32))
		}
	case uint64:
		{
			res = float64(value.(uint64))
		}
	case uintptr:
		{
			res = float64(value.(uintptr))
		}

	case float32:
		{
			res = float64(value.(float32))
		}
	case float64:
		{
			res = float64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToFloat64(valueStr)
		}
	}
	return
}
