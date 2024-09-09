package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {
	fmt.Println("Введите выражение:")
	var input string
	fmt.Scanln(&input)

	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		fmt.Println("Ошибка! Неизвестный оператор.")
		os.Exit(1)
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		fmt.Println("Ошибка! Неверный формат ввода.")
		os.Exit(1)
	}

	aStr := parts[0]
	bStr := parts[1]

	isRoman := false

	if _, ok := romanToArabic[aStr]; ok {
		isRoman = true
		if _, ok := romanToArabic[bStr]; !ok {
			fmt.Println("Ошибка! Оба числа должны быть римскими или арабскими.")
			os.Exit(1)
		}
	} else {
		a, err := strconv.Atoi(aStr)
		if err != nil || a < 1 || a > 10 {
			fmt.Println("Ошибка! Числа должны быть в диапазоене от 1 до 10.")
			os.Exit(1)
		}
		b, err := strconv.Atoi(bStr)
		if err != nil || b < 1 || b > 10 {
			fmt.Println("Ошибка! Числа должны быть в диапазоне от 1 до 10.")
			os.Exit(1)
		}
	}

	var a, b int
	if isRoman {
		a = romanToInt(aStr)
		b = romanToInt(bStr)
	} else {
		a, _ = strconv.Atoi(aStr)
		b, _ = strconv.Atoi(bStr)
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка! Деление на ноль.")
			os.Exit(1)
		}
		result = a / b
	default:
		fmt.Println("Ошибка! Неизвестный оператор.")
		os.Exit(1)
	}

	if isRoman {
		if result <= 0 {
			panic("Результат для римских чисел должен быть больше нуля.") //Исправил
		}
		fmt.Println("Результат:", intToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}

func romanToInt(roman string) int {
	return romanToArabic[roman]
}

func intToRoman(num int) string {
	if num <= 0 {
		panic("Результат римских чисел должен быть больше нуля.") //Исправил
	}

	roman := ""
	for _, entry := range arabicToRoman {
		for num >= entry.Value {
			roman += entry.Symbol
			num -= entry.Value
		}
	}
	return roman
}
