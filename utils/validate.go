package utils

import (
	"fmt"
	"github.com/noaway/dateparse"
	"reflect"
	"regexp"
	"time"

	"github.com/zxbit2011/ant/utils/convert"
	"strings"
)

const (
	regularMobile      = `^((\+86)|(86))?(13\d|15[^4\D]|17[13678]|18\d)\d{8}|170[^346\D]\d{7}$`
	regularMobilePhone = `^(1)\d{10}$`
	regularEmail       = `^([\w-_]+(?:\.[\w-_]+)*)@((?:[a-z0-9]+(?:-[a-zA-Z0-9]+)*)+\.[a-z]{2,6})$`
	regularIPv4        = `^(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)$`
	//任意实数，包括带正号的正数
	regularNumber = `^(-|\+)?(([1-9]\d*(.\d*)?)|0(.\d*)?)$`
)

//检查数据类型（基于字符串）
func CheckType(value interface{}, expectType string) bool {
	valueTypeString := reflect.TypeOf(value).String()
	return CheckRegexp(valueTypeString, `^`+expectType+`((32)|(64))?$`)
}

//检查int数据的区间（开区间）
func CheckIntRange(value interface{}, min, max int) bool {
	if val, err := convert.ToInt(value); err != nil {
		return false
	} else {
		return min <= val && val <= max
	}
}

//检查float数据的区间(开区间)
func CheckFloat64Range(value interface{}, min, max float64) bool {
	if !CheckRealNumber(value) {
		return false
	}

	if val, err := convert.ToFloat64(value); err != nil {
		return false
	} else {
		return min <= val && val <= max
	}
}

//根据正则检查字符串，如果传入的不是string类型，则使用Sprintf处理成字符串再校验。
func CheckRegexp(value interface{}, rex string) bool {
	val, ok := value.(string)
	if !ok {
		val = convert.MustString(value)
	}
	reg := regexp.MustCompile(rex)
	return reg.MatchString(val)
}

//检查email
func CheckEmail(value interface{}) bool {
	return CheckRegexp(value, regularEmail)
}

//检查手机
func CheckMobile(value interface{}) bool {
	return CheckRegexp(value, regularMobile)
}

// 合法的IPV4
func CheckIPv4(value interface{}) bool {
	return CheckRegexp(value, regularIPv4)
}

//检查传入值是否是实数（包括表示实数的字符串）
func CheckRealNumber(value interface{}) bool {
	val := convert.MustString(value)
	return CheckRegexp(val, regularNumber)
}

//检查传入的值的长度，仅支持string、slice或者map
// 计算String的时候，以正则实现，中文字符算长度1
func CheckLen(value interface{}, length int) bool {
	switch value.(type) {
	case string:
		return CheckRegexp(value, fmt.Sprintf("^.{%d}$", length))
	default:
		{
			refValue := reflect.ValueOf(value)
			switch refValue.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				{
					return refValue.Len() == length
				}
			}
		}
	}

	return false
}

//检查传入的数值的下限,基本上支持各种类型
func CheckMin(value interface{}, min float64) bool {
	switch value.(type) {
	case string:
		{
			res, err := convert.ToFloat64(value)
			if err == nil {
				return res > float64(min)
			}
		}
	case int, int8, int16, int32, int64:
		{
			return convert.MustInt64(value) >= int64(min)
		}
	case uint, uint8, uint16, uint32, uint64:
		{
			return convert.MustUint64(value) >= uint64(min)
		}
	case float32, float64:
		{
			return convert.MustFloat64(value) >= float64(min)
		}
	default:
		{
			return CheckMin(convert.MustString(value), min)
		}
	}
	return false
}

//检查传入的数值的上限,基本上支持各种类型
func CheckMax(value interface{}, max float64) bool {
	switch value.(type) {
	case string:
		{
			res, err := convert.ToFloat64(value)
			if err == nil {
				return res <= float64(max)
			}
		}
	case int, int8, int16, int32, int64:
		{
			return convert.MustInt64(value) <= int64(max)
		}
	case uint, uint8, uint16, uint32, uint64:
		{
			return convert.MustUint64(value) <= uint64(max)
		}
	case float32, float64:
		{
			return convert.MustFloat64(value) <= float64(max)
		}
	default:
		{
			return CheckMin(convert.MustString(value), max)
		}
	}
	return false
}

// 检查传入的值的长度，仅支持string、slice或者map
// 计算String的时候，以正则实现，中文字符算长度1
func CheckMaxSize(value interface{}, maxSize int) bool {
	switch value.(type) {
	case string:
		{
			return CheckRegexp(value, fmt.Sprintf("^.{0,%d}$", maxSize))
		}
	default:
		{
			refValue := reflect.ValueOf(value)
			switch refValue.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				{
					return refValue.Len() <= maxSize
				}
			}
		}
	}

	return false
}

//检查传入的值的长度，仅支持string、slice或者map
// 计算String的时候，以正则实现，中文字符算长度1
func CheckMinSize(value interface{}, minSize int) bool {
	switch value.(type) {
	case string:
		{
			return CheckRegexp(value, fmt.Sprintf("^.{%d,}$", minSize))
		}
	default:
		{
			refValue := reflect.ValueOf(value)
			switch refValue.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				{
					return refValue.Len() >= minSize
				}
			}
		}
	}

	return false
}

func IsMobile(s string) bool {
	return CheckRegexp(s, regularMobilePhone)
}

// 检查字符串是否代表一个合法的bool值
// 注意返回值仅表示是不是布尔值，并不表示布尔值本身的含义
func IsValidBoolean(s string) bool {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True", "0", "f", "F", "false", "FALSE", "False":
		return true
	}
	return false
}

func IsValidDateTime(s string) bool {
	t, err := dateparse.ParseAny(s)
	if err != nil || t.Year() < 1971 {
		return false
	}
	return true
}

func IsValidTime(s string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return false
	}
	return true
}

func IsValidDate(s string) bool {
	s = strings.Replace(s, "/", "-", -1)
	s = strings.Replace(s, ".", "-", -1)
	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		return false
	}
	return true
}

// 检查字符串是不是一个合法的数字
func IsValidNumber(s string) bool {
	// This function implements the JSON numbers grammar.
	// See https://tools.ietf.org/html/rfc7159#section-6
	// and http://json.org/number.gif

	if s == "" {
		return false
	}

	// Optional -
	if s[0] == '-' {
		s = s[1:]
		if s == "" {
			return false
		}
	}

	// Digits
	switch {
	default:
		return false

	case s[0] == '0':
		s = s[1:]

	case '1' <= s[0] && s[0] <= '9':
		s = s[1:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// . followed by 1 or more digits.
	if len(s) >= 2 && s[0] == '.' && '0' <= s[1] && s[1] <= '9' {
		s = s[2:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// e or E followed by an optional - or + and
	// 1 or more digits.
	if len(s) >= 2 && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
			if s == "" {
				return false
			}
		}
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// Make sure we are at the end.
	return s == ""
}

// 传入参数是否一个数组或者切片
func IsArray(value interface{}) bool {
	refValue := reflect.ValueOf(value)
	switch refValue.Kind() {
	case reflect.Slice, reflect.Array:
		{
			return true
		}
	default:
		return false
	}

}

func IsUrl(url string) bool {
	reStr := "http(s)?://([\\w-]+\\.)+[\\w-]+(/[\\w- ./?%&=]*)?"
	flog, _ := regexp.Match(reStr, []byte(url))
	return flog
}

func IsWechatBrowser(userAgent string) bool {
	if strings.Contains(strings.ToLower(userAgent), "micromessenger") {
		return true
	}
	return false
}
