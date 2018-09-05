package main

import (
	"fmt"
	"math"
)

// 判断是否是素数
func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	var a int
	var b int
	// 任意输出两个数之间的素数

	fmt.Scanf("%d%d", &a, &b)
	// fmt.Println(a, b)
	// 如果a大于b，互换
	if int(a) > int(b) {
		a, b = b, a
	}

	for i := a; i < b; i++ {

		if a < 2 {
			fmt.Printf("素数是大于1的自然数除了1和他本身不能有其他整除")
		}
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}
