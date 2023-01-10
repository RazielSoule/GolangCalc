// Калькулятор Golang
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for { //В тз не было цикличности программы но так удобнее
		fmt.Println("Введите два операнда и операцию для вычисления в формате: \"a + b\"")
		var rawValue1, operation, rawValue2 string
		var isRoman1, isRoman2 bool
		var result int
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')
		inputSep := strings.Split(input[:len(input)-2], " ")
		if len(inputSep) != 3 {
			fmt.Println("Ошибка - непримелимый формат ввода")
			os.Exit(6)
		}
		rawValue1 = inputSep[0]
		operation = inputSep[1]
		rawValue2 = inputSep[2]

		//fmt.Scan(&rawValue1, &operation, &rawValue2)
		//fmt.Println(rawValue1, operation, rawValue2)
		//fmt.Println(len(rawValue1), len(operation), len(rawValue2))

		value1, err := strconv.Atoi(rawValue1)
		if err != nil {
			isRoman1 = true
			//fmt.Println("Первая римская")
		}
		value2, err := strconv.Atoi(rawValue2)
		if err != nil {
			isRoman2 = true
			//fmt.Println("Вторая римская")
		}

		if isRoman1 && isRoman2 {
			// //Если оба операнда - римские
			// value1 = toArabicConvert(strings.ToUpper(rawValue1))
			// value2 = toArabicConvert(strings.ToUpper(rawValue2))
			value1 = romanToInt(strings.ToUpper(rawValue1))
			value2 = romanToInt(strings.ToUpper(rawValue2))
			result = calculate(value1, value2, operation)
			if result > 0 {
				fmt.Println(toRomanConvert(result))
				//fmt.Println(result)
			} else {
				fmt.Println("Ошибка - результат операции с римскими цифрами меньше I")
				os.Exit(5)
			}
		} else if !isRoman1 && !isRoman2 {
			//Если оба операнда - арабские
			if value1 <= 10 && value1 > 0 && value2 <= 10 && value2 > 0 {
				result = calculate(value1, value2, operation)
				fmt.Println(result)
			} else {
				fmt.Println("Ошибка - значение на входе не в диапазоне 1-10 включительно")
				os.Exit(4)
			}
		} else {
			fmt.Println("Ошибка - введены операнды разного типа")
			os.Exit(3)
		}
	}
}

// Вычисляет результат, он тут шишка важная, ага
func calculate(a, b int, operation string) (result int) {
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("Ошибка - введён неприемлимый оператор")
		os.Exit(2)
	}
	return
}

// Уверен, это ужасный велосипед, но он конвертирует арабские цифры в римские (корректные значение до 400)
func toRomanConvert(arabic int) (romanOut string) {
	counter := 0
	romanOut = ""
	for arabic != 0 {
		switch {
		case arabic >= 100:
			counter = arabic / 100
			arabic = arabic % 100
			for ; counter != 0; counter-- {
				romanOut += "C"
			}
		case arabic >= 90:
			arabic = arabic % 90
			romanOut += "XC"
		case arabic >= 50:
			arabic = arabic % 50
			romanOut += "L"
		case arabic >= 10:
			counter = arabic / 10
			arabic = arabic % 10
			if counter == 4 {
				romanOut += "XL"
				counter = 0
			}
			for ; counter != 0; counter-- {
				romanOut += "X"
			}
		case arabic == 9:
			arabic = 0
			romanOut += "IX"
		case arabic >= 5:
			arabic = arabic % 5
			romanOut += "V"
		case arabic >= 1:
			counter = arabic
			arabic = 0
			if counter == 4 {
				romanOut += "IV"
				counter = 0
			}
			for ; counter != 0; counter-- {
				romanOut += "I"
			}
		}
	}
	return
}

// Конвертирует введённую римскую в арабскую (до XI)
func toArabicConvert(roman string) (arabicOut int) {
	switch roman {
	case "I":
		arabicOut = 1
	case "II":
		arabicOut = 2
	case "III":
		arabicOut = 3
	case "IV":
		arabicOut = 4
	case "V":
		arabicOut = 5
	case "VI":
		arabicOut = 6
	case "VII":
		arabicOut = 7
	case "VIII":
		arabicOut = 8
	case "IX":
		arabicOut = 9
	case "X":
		arabicOut = 10
	default:
		fmt.Println("Ошибка - введённый операнд не соответсвует арабскому или римскому числу или превышает X")
		os.Exit(1)
	}
	return
}

func romanToInt(s string) int { // Взято из сети для проверки
	characterMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return characterMap[s[0]]
	}

	sum := characterMap[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		if characterMap[s[i]] < characterMap[s[i+1]] {
			sum -= characterMap[s[i]]
		} else {
			sum += characterMap[s[i]]
		}
	}

	return sum
}
