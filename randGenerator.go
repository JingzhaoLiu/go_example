package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randGenerator(n int) chan int {

	
	out := make(chan int)
	go func(x int) {
		for {
			out <- rand.Intn(x)
		}
	}(n)
	return out
	// out := rand.Intn(n)
	// return out
}
func main() {
	// 生成随机数作为一个服务

	// 从服务中读取随机数并打印
	t1 := time.Now()
	for i := 0; i < 100000; i++ {

		rand_service_handler := randGenerator(100)

		fmt.Printf("%d\n", <-rand_service_handler)

	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	// 测试 不知道为什么使用通道比不使用通道速度慢

}
