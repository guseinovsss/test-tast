package main

import (
	"fmt"
	"strconv"
)

func arabic(a, b int, c string) {

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

func testForArabic(a1, b1, c string) int {
	var a, err = strconv.Atoi(a1)
	var b, err1 = strconv.Atoi(b1)
	if nil != err || nil != err1 {
		fmt.Println("Оба должны быть числом")
		return 0
	}

	if (0 < a && a < 11) && (0 < b && b < 11) {
		arabic(a, b, c)
	} else {
		fmt.Println("Числа должны быть больше от 1 до 10 включительно")
		return 0
	}

	return 1
}

func main() {
	var a, b, c string
	fmt.Scan(&a, &c, &b)

	testForArabic(a, b, c)

}
