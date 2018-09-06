package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* 对于Random的使用，在业务中使用频率是非常高的，本文就小结下常用的方法：

在Golang中，有两个包提供了rand，分别为 "math/rand" 和 "crypto/rand",  对应两种应用场景。

一、"math/rand" 包实现了伪随机数生成器。也就是生成 整形和浮点型。

　　 该包中根据生成伪随机数是是否有种子(可以理解为初始化伪随机数)，可以分为两类：

　　1、有种子。通常以时钟，输入输出等特殊节点作为参数，初始化。该类型生成的随机数相比无种子时重复概率较低。

　　2、无种子。可以理解为此时种子为1， Seek(1)。

常用的方法有：(以有种子的为例，无种子的直接通过 rand 报名调用对应的方法)

　　　　1> 按类型随机类：

		func (r *Rand) Int() int
		func (r *Rand) Int31() int32
		func (r *Rand) Int63() int64
		func (r *Rand) Uint32() uint32
		func (r *Rand) Float32() float32  // 返回一个取值范围在[0.0, 1.0)的伪随机float32值
		func (r *Rand) Float64() float64  // 返回一个取值范围在[0.0, 1.0)的伪随机float64值
　　　　2>指定随机范围类：

		func (r *Rand) Intn(n int) int
		func (r *Rand) Int31n(n int32) int32
		func (r *Rand) Int63n(n int64) int64

*/

func init() {
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}
func main() {
	// 生成随机数
	// for i := 0; i < 10; i++ {
	// 	a := rand.Int()
	// 	fmt.Println(a)
	// }
	// 生成[0,10)随机数
	// for i := 0; i < 10; i++ {
	// 	a := rand.Intn(10)
	// 	fmt.Println(a)
	// }

	// 生成[0,100)随机数
	// for i := 0; i < 50; i++ {
	// 	a := rand.Intn(100)
	// 	fmt.Println(a)
	// }
	// 生成[1,10)随机数
	// for i := 0; i < 100; i++ {
	// 	a := rand.Intn(9) + 1
	// 	fmt.Println(a)
	// }
	// 生成[0.0, 1.0)随机float32值
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
	// 生成[0.0, 1.0)随机float64值
	for i := 0; i < 10; i++ {
		a := rand.Float64()
		fmt.Println(a)
	}

	// 2^32-1(+4 294 967 295)
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Int31())

	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// 2^63-1(+9,223,372,036,854,775,807 )
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int63())

	}

	// 对于需要随机指定位数的，当位数不够是，可以通过前边补0达到长度一致
	// 随机产生5位长度伪随机数
	for i := 0; i < 100; i++ {
		fmt.Printf("%.5d ", rand.Int31()%100)
		fmt.Printf("%.5d ", rand.Int31()%10000)
	}

}

/*二、”crypto/rand“ 包实现了用于加解密的更安全的随机数生成器。

　　该包中常用的是 func Read(b []byte) (n int, err error) 这个方法， 将随机的byte值填充到b 数组中，以供b使用。
*/

// import (
// 	"crypto/rand"
// 	"fmt"
// )

// func main() {
// 	b := make([]byte, 20)
// 	fmt.Println(b) //

// 	_, err := rand.Read(b)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println(b)
// }
