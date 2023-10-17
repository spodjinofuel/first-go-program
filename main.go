package main

import (
	"fmt"
	"strings"
)

func exist(numbers []string, number string) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

func calc(a, b string, ch string) interface{} {
	isArab := func(number string) bool {
		arabnum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		return exist(arabnum, number)
	}

	isRom := func(number string) bool {
		romnum := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
		return exist(romnum, number)
	}

	isArabA := isArab(a)
	isArabB := isArab(b)
	isRomA := isRom(a)
	isRomB := isRom(b)

	toRom := func(number int) string {
		romnums := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C"}
		var result strings.Builder
		keys := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
		for _, value := range keys {
			for number >= value {
				result.WriteString(romnums[value])
				number -= value
			}
		}
		return result.String()
	}

	toArab := func(number string) int {
		romnums := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100}
		result := 0
		i := 0
		for i < len(number) {
			if i+1 < len(number) && romnums[number[i:i+2]] != 0 {
				result += romnums[number[i:i+2]]
				i += 2
			} else {
				result += romnums[string(number[i])]
				i += 1
			}
		}
		return result
	}

	if isArabA && isArabB {
		var aInt, bInt int
		fmt.Sscanf(a, "%d", &aInt)
		fmt.Sscanf(b, "%d", &bInt)
		switch ch {
		case "+":
			return aInt + bInt
		case "-":
			return aInt - bInt
		case "*":
			return aInt * bInt
		case "/":
			return aInt / bInt
		default:
			return "Калькулятор не поддерживает введенную операцию"
		}
	} else if isRomA && isRomB {
		aInt := toArab(a)
		bInt := toArab(b)
		switch ch {
		case "+":
			return toRom(aInt + bInt)
		case "-":
			return toRom(aInt - bInt)
		case "*":
			return toRom(aInt * bInt)
		case "/":
			return toRom(aInt / bInt)
		default:
			return "Калькулятор не поддерживает введенную операцию"
		}
	} else if isRomA && isArabB || isArabA && isRomB {
		return "Введенные числа из разных систем счисления"
	} else {
		return "Неверно введенная строка"
	}
}

func main() {
	var a, b, ch string
	fmt.Scan(&a, &ch, &b)
	result := calc(a, b, ch)
	fmt.Println(result)
}
