package main

import (
	"fmt"
	"strconv"
)

func firstTask() {
	var num string
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	reversedNum := reverseNumber(num)
	fmt.Printf("Перевёрнутое число: %s\n", reversedNum)
}

func reverseNumber(numStr string) string {
	runes := []rune(numStr)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversedStr := string(runes)
	return reversedStr
}

func secondTask() {
	var a, b, c float64
	fmt.Print("Введите три стороны треугольника: ")
	fmt.Scan(&a, &b, &c)

	if isRightTriangle(a, b, c) {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

func isRightTriangle(a, b, c float64) bool {
	a2, b2, c2 := a*a, b*b, c*c

	return a2+b2 == c2 || a2+c2 == b2 || b2+c2 == a2
}

func thirdTask() {
	var k int
	fmt.Print("Введите количество секунд: ")
	fmt.Scan(&k)

	h := k / 3600
	m := (k % 3600) / 60

	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}

func fourthTask() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	result := squareDigits(num)
	fmt.Printf("Результат: %s\n", result)
}

func squareDigits(num int) string {
	numStr := strconv.Itoa(num)
	var result string

	for _, digit := range numStr {
		n, _ := strconv.Atoi(string(digit))
		squared := n * n
		result += strconv.Itoa(squared)
	}

	return result
}

func fifthTask() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scan(&input)

	result := insertAsterisks(input)
	fmt.Printf("Результат: %s\n", result)
}

func insertAsterisks(s string) string {
	if len(s) <= 1 {
		return s
	}

	var result string
	result = string(s[0])

	for i := 1; i < len(s); i++ {
		result += "*" + string(s[i])
	}

	return result
}

func main() {
	fourthTask()
}
