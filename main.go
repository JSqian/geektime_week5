package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数据保存
var LimitQueue map[string][]int64

func LimitHandle(name string, count int, timeWindow int64) bool {
	currTime := time.Now().Unix()

	//创建
	if _, ok := LimitQueue[name]; !ok {
		LimitQueue[name] = make([]int64, 0)
	}

	//队列未满
	if len(LimitQueue[name]) < count {
		LimitQueue[name] = append(LimitQueue[name], currTime)
		return true
	}

	//队列满了
	//找到最早的那个时间
	firstTime := LimitQueue[name][0]
	//小于窗口的限定时间，说明还没过期，此次请求不允许通过
	if currTime-firstTime <= timeWindow {
		return false
	}

	//校验通过，将本次请求添加到队列中
	LimitQueue[name] = LimitQueue[name][1:]
	LimitQueue[name] = append(LimitQueue[name], currTime)

	return true
}

func init() {
	LimitQueue = make(map[string][]int64)
}

// 使用示例
func main() {
	//timeWindow秒内，可以容纳count次请求
	req := "test1"
	var timeWindow int64 = 10
	var count int = 5

	for i := 0; i < 20; i++ {
		res := LimitHandle(req, count, timeWindow)
		fmt.Printf("第%d次请求 是否通行:%t\n", i+1, res)
		sec := rand.Intn(5)
		time.Sleep(time.Duration(sec) * time.Second)
	}
}
