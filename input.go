package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ExitRequestError запрошен выход из программы
type ExitRequestError struct{}

func (e ExitRequestError) Error() string {
	return "пользователь захотел выйти из программы"
}

// NotValidNumberError неверный ввод номера карты
type NotValidNumberError struct{}

func (e NotValidNumberError) Error() string {
	return "Неверный номер карты, повторите ввод."
}

// userInput получение пользовательского ввода
func userInput() string {
	fmt.Print("Введите номер кредитной карты (для выхода сразу нажмите ВВОД): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	processed := strings.TrimSpace(input)

	return processed
}

// isValidInput проверка пользовательского ввода
func isValidInput(cardNumber string, minLength, maxLength int) bool {
	length := len(cardNumber)

	if length < minLength || length > maxLength {
		return false
	}

	for _, s := range cardNumber {
		if !unicode.IsDigit(s) {
			return false
		}
	}

	return true
}

// getCardNumber получить от пользователя корректный номер карты
func getCardNumber(minLength, maxLength int) (string, error) {
	input := userInput()
	if len(input) == 0 {
		return "", ExitRequestError{}
	}
	isValid := isValidInput(input, minLength, maxLength)
	if isValid {
		return input, nil
	}
	return "", NotValidNumberError{}
}
