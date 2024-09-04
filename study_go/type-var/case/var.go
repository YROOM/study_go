package main

import "fmt"

func VarDeclareCase() {

	/**
	值类型
	*/
	var i int

	var j int = 100

	var f = 100.123456

	arr := [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7}

	var arr2 [5]int

	arr2[2] = 4
	arr2[3] = 5

	for i2, i3 := range arr {
		fmt.Println(i2, i3)
	}

	for i2, i3 := range arr1 {
		fmt.Println(i2, i3)
	}

	fmt.Println(f, j, i)

	// 指针类型，用来表示变量的地址

	//	var intPtr *int

	var i1 = 100
	f1(&i1)
	fmt.Println(i1)

}

func f1(i *int) {
	*i = *i + 1
}
