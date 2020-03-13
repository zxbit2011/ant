package convert

// 强制转换为int，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt(value interface{}) (res int) {
	if res, err := ToInt(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为int8，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt8(value interface{}) (res int8) {
	if res, err := ToInt8(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为int16，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt16(value interface{}) (res int16) {
	if res, err := ToInt16(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为int32，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt32(value interface{}) (res int32) {
	if res, err := ToInt32(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为int64，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt64(value interface{}) (res int64) {
	if res, err := ToInt64(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为uint，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint；空字符串会转化为0.
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint(value interface{}) (res uint) {
	if res, err := ToUint(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为uint8，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint8(value interface{}) (res uint8) {
	if res, err := ToUint8(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为uint16，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint16(value interface{}) (res uint16) {
	if res, err := ToUint16(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为uint32，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint32(value interface{}) (res uint32) {
	if res, err := ToUint32(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// 强制转换为uint64，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint64(value interface{}) (res uint64) {
	if res, err := ToUint64(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

func MustFloat64(value interface{}) (res float64) {
	if res, err := ToFloat64(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

func MustFloat32(value interface{}) (res float32) {
	if res, err := ToFloat32(value); err == nil {
		return res
	} else {
		panic(err)
	}
}
