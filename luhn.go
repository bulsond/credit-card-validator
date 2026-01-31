package main

import (
	"strconv"
)

// validateLuhn проверка строки номера карты по алгоритму Луна
func validateLuhn(cardNumber string) bool {

	// Шаг 1: Преобразовать входную строку в массив цифр
	digits := getDigits(cardNumber)

	// Шаг 2: Применить преобразование Луна
	for i := len(digits) - 1; i >= 0; i-- {
		if i%2 != 0 {
			continue
		}
		n := digits[i] * 2
		if n > 9 {
			n -= 9
		}
		digits[i] = n
	}

	// Шаг 3: Вычисляем сумму элементов полученного массива
	var sum uint8
	for _, v := range digits {
		sum += v
	}

	// Шаг 4: Проверяем, что полученная сумма делится на 10 без остатка
	result := (sum % 10) == 0

	return result
}

func getDigits(cardNumber string) []uint8 {
	var digits []uint8
	for _, s := range cardNumber {
		digit, _ := strconv.Atoi(string(s))
		digits = append(digits, uint8(digit))
	}
	return digits
}
