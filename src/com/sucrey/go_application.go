package main

import (
	"com/sucrey/service"
	"fmt"
	"math/rand"
	"unicode/utf8"
)

var _ = service.Test

func init() {
	fmt.Println("go_application has been init")
}

type Employee struct {
	firstName  string
	secondName string
	age        int
}

func main() {
test:
	for i := 1; i < 10; i++ {
		if i == 5 {
			break test
		}
		fmt.Println(i)
	}
	switch num := rand.Int(); { // expression is omitted
	case num < 50:
		fmt.Println("num is greater than 0 and less than 50")
		fallthrough
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

	a := [...]float64{22.2, 23, 55, 67.1}
	b := make([]float64, 5, 8)
	var c = []float64{2}
	ints := append(c, b...)
	fmt.Println(ints)
	for index, value := range a {
		fmt.Println(index, value)
		fmt.Println(b)
		ints := append(b, 9, 6, 2, 2, 1)
		fmt.Println(ints)
		fmt.Println(cap(ints))
	}

	var d map[string]int
	if d == nil {
		fmt.Println("d is nil ,init map")
		d = make(map[string]int)
	}
	d["sb"] = 23
	fmt.Println(d)

	e := map[string]string{
		"test": "2222",
		"yy":   "123",
		"rr":   "2111"}

	f := map[string]string{
		"test": "2222",
		"yy":   "123",
		"rr":   "2111"}
	i := equals(e, f)
	fmt.Println(i)

	_, ok := f["yy"]
	fmt.Println(ok)

	j := "123"
	fmt.Printf("%x", j[0])
	fmt.Println(j[1])
	fmt.Println(j[2])

	runes := []rune(j)
	fmt.Printf("%c\n", runes[1])
	n := utf8.RuneCountInString(j)
	fmt.Println(n)

	em1 := Employee{"asd", "asd", 222}
	fmt.Println(em1)
	em1.getEmployeeInfo()

}

func (e Employee) getEmployeeInfo() {
	fmt.Println(e.age, e.firstName, e.secondName)
}

/*
	判断两个map是否相等
*/
func equals(source map[string]string, target map[string]string) bool {
	if len(source) != len(target) {
		return false
	}
	for k, v := range source {
		_, hasKey := target[k]
		if !hasKey {
			fmt.Printf("don't have the key: %s \n", k)
			return false
		}
		if v != target[k] {
			fmt.Println("value is not equal")
			return false
		}
	}
	return true
}
