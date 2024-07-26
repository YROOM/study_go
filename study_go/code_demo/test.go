package main

func main() {
	var num int = 9

	var ptr *int

	ptr = &num

	*ptr = 10

	println("num=", num)
}
