package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RandID struct {
	mutex         sync.Mutex
	lastTimestamp int64 //最后时间戳
	stamp         map[int]int
	errCount      int // 尝试错误次数
}

var randID *RandID
var nextId int

func init() {
	randID = NewRandID()
}
func NewRandID() *RandID {
	randID = &RandID{}
	randID.mutex = sync.Mutex{}
	return randID
}
func (id *RandID) NextId() int64 {
	id.mutex.Lock()
	defer id.mutex.Unlock()
	return id.nextId()
}

func (id *RandID) nextId() int64 {
	newNextId := rand.Intn(999)
	//一毫秒产生一个ID保障唯一
	mill := time.Now().UnixNano() / 1e6
	if mill == id.lastTimestamp {
		//判断是否重复，重复尝试次数100次
		if id.stamp[newNextId] == 1 && id.errCount <= 100 {
			//重复重新获取
			id.errCount++
			return id.nextId()
		}
		id.stamp[newNextId] = 1
	} else {
		id.stamp = map[int]int{}
		id.stamp[newNextId] = 1
	}
	id.errCount = 0
	nextId = newNextId
	id.lastTimestamp = time.Now().UnixNano() / 1e6
	uuid, err := strconv.ParseInt(fmt.Sprintf("%v%v", time.Now().UnixNano()/1e6, fmt.Sprintf("%03d", nextId)), 10, 64)
	if err != nil {
		panic(err)
	}
	return uuid
}

func IDInt() int64 {
	return randID.NextId()
}

func ID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return strings.Replace(fmt.Sprintf("%s", id), "-", "", -1)
}
