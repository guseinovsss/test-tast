package main

import (
	"fmt"
	"strconv"
)

func arabic(a1, b1, c string) {
	var a, err = strconv.Atoi(a1)
	var b, err1 = strconv.Atoi(b1)
	if nil == err {
	}
	if nil == err1 {
	}
	var r int = 0
	switch c {
	case "+":
		fmt.Println(a + b)
		r++
	case "-":
		fmt.Println(a - b)
		r++
	case "/":
		fmt.Println(a / b)
		r++
	case "*":
		fmt.Println(a * b)
		r++
	}
	if r == 0 {
		fmt.Println("Такого действия нет")
	}
}

func main() {
	var a, b, c string
	fmt.Scan(&a, &c, &b)

	arabic(a, b, c)

}
