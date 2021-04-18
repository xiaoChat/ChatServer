package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"xiaochat/app/db"
	"xiaochat/app/model"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	// lock.Lock() // 加锁
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	// lock.Unlock() // 解锁
	wg.Done()
}
func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
func chanq() {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func rec(c chan int) {
	res := <-c
	fmt.Println("发送成功" + strconv.Itoa(res))
}

func hello(start int) {
	fmt.Println(strconv.Itoa(start) + "===>" + time.Now().Format(time.RFC3339))
}

func dbtest() {
	db, err := db.GetYexk()
	if err != nil {
		panic("mysql connect error!")
	}

	a := 10
	b := &a
	c := a
	a = 20
	fmt.Printf("a:%d, b:%d, c:%d \n", a, *b, c)

	go hello(1)
	go hello(2)
	hello(3)
	hello(4)

	_ = db.Model(&model.User{}).AutoMigrate(&model.User{})
	db.Model(&model.User{}).Create(&model.User{
		Name:     "yexk1",
		Email:    "yex1k@yexk.cn",
		Age:      10,
		Birthday: time.Now(),
	})
}
