package service

import "fmt"

func Test(a int) int {
	return a + 1
}

func init() {
	fmt.Println("go_service has been init")
}
