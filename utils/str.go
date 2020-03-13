package utils

import (
	"crypto/sha1"
	"fmt"
	"path"
	"regexp"
	"runtime"
	"strings"
)

func SubStrByByte(str string, length int) string {
	bt := []byte(str)
	if len(bt) <= length {
		return str
	}
	bs := bt[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

func SubStrByByteInChar(str string, length int) string {
	s := SubStrByByte(str, length-8)
	if s == str {
		return str
	}
	return SubStrByByte(str, length-8) + ".."
}

func MobileReplaceRepl(str string) string {
	re, _ := regexp.Compile("(\\d{3})(\\d{4})(\\d{4})")
	return re.ReplaceAllString(str, "$1****$3")
}

func Contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func ContainsInt64(strs []int64, str int64) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

//去除前后所有空格、空字符串、制表符
func Trim(str string) string {
	if str == "" {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(str, string('\uFEFF')))
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//清除空格
	src = strings.Replace(src, "&nbsp;", "", -1)

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

func GetFuncName() string {
	funcName, _, _, ok := runtime.Caller(1)
	if ok {
		_, filename := path.Split(runtime.FuncForPC(funcName).Name())
		return filename
	}
	return ""
}

func GetHash(str string) string {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(str))
	return string(Sha1Inst.Sum([]byte("")))
}

func DeleteRepeat(list []string) []string {
	data := make(map[string]string)
	if len(list) <= 0 {
		return nil
	}
	for _, v := range list {
		data[v] = "true"
	}
	var datas []string
	for k := range data {
		if k == "" {
			continue
		}
		datas = append(datas, k)
	}
	return datas
}

func Reverse(l []string) {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		fmt.Println(i, "<=>", li)
		l[i], l[li] = l[li], l[i]
	}
}
