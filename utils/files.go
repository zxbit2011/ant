package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Read(path string) string {
	return string(ReadByte(path))
}

func ReadByte(path string) []byte {
	fi, err := os.Open(path)
	if err != nil {
		println("Files Logs Error:" + err.Error())
		panic(err)
	}
	defer fi.Close()
	buf, err := ioutil.ReadAll(fi)
	if err != nil {
		println("Files Logs Error:" + err.Error())
		panic(err)
	}
	return buf
}

//系统分隔符
func Separator() string {
	var path string
	//前边的判断是否是系统的分隔符
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	return path
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

//当前目录
func CurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建多级文件夹
func MkdirAll(path string) error {
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = strings.Replace(path, "/", "\\", -1)
	}
	flog, err := PathExists(path)
	if err != nil {
		return err
	}
	if flog {
		return nil
	}
	err2 := os.MkdirAll(path, os.ModePerm)

	if err2 != nil {
		return err2
	}
	return nil
}

func CreateFile(path string) (*os.File, error) {
	i := strings.LastIndex(path, "/")
	dir := string([]rune(path)[0:i])
	mErr := MkdirAll(dir)
	if mErr != nil {
		return nil, mErr
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func JsonPath(params ...string) string {
	cp := CurrentDirectory()
	for _, p := range params {
		cp += Separator() + p
	}
	return cp
}

func JsonParentPath(params ...string) string {
	cp := getParentDirectory(CurrentDirectory())
	for _, p := range params {
		cp += Separator() + p
	}
	return cp
}

// 判断是否是文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 判断是否是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
