package main

import (
	"fmt"
	p2 "study_go/lan-structure/packgae1"
)

// 每个包在引用的时候都会调用
// 初始化函数  golang 每个包的引用会有限调用该函数
func init() {
	fmt.Println("init function")
}

func main() {

	var i = 10
	p2.F1()

	fmt.Println("调用了 lan-structure main")
	fmt.Println(fmt.Sprintf("打印参数i: %d", i))

}

// 先执行init 函数， 在执行main函数
