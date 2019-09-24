package main

import ""
import "fmt"

const aaa = 1

const test1 = "123"


func main() {
	var b = aaa

	fmt.Printf("type %T value %v", b, b)
	i, j := test(2)
	fmt.Println()
	fmt.Println(i,j)
}

func test(a int) (b ,c int) {
	b = a + 1
	c = a + 2
	return
}
