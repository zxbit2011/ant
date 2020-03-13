package utils

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestGetValid(t *testing.T) {
	str := "zxtestcv，Valid[Email;Min(120);]()][fdwd]"
	var reg = regexp.MustCompile(`Valid\[(.*?)]`)
	params := reg.FindStringSubmatch(str)
	for _, param := range params {
		t.Log(strings.TrimRight(param, ";"))
	}
}
func TestGetValid3(t *testing.T) {
	str := "zxtestcv，Valid[Email;Min(120);Range[1,20]]()][fdwd]"
	var reg = regexp.MustCompile(`Valid\[([^[]] | .*?)*]`)
	params := reg.FindStringSubmatch(str)
	for _, param := range params {
		t.Log(param)
	}
}

//传入[]byte，返回[]byte
func TestGetValid2(t *testing.T) {
	str := "ab001234hah120210a880218end"
	reg := regexp.MustCompile("\\d{6}") //六位连续的数字
	fmt.Println("------Find------")
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	fmt.Println(string(data))

	fmt.Println("------FindAll------")
	//返回str中所有匹配reg的字符串
	//第二个参数表示最多返回的个数，传-1表示返回所有结果
	dataSlice := reg.FindAll([]byte(str), -1)
	for _, v := range dataSlice {
		fmt.Println(string(v))
	}
}
