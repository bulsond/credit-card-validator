package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Bank данные о банке
type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

// loadBankData загрузка из файла данных о банках
func loadBankData(path string) ([]Bank, error) {
	banks := []Bank{}

	file, err := os.Open(path)
	if err != nil {
		return banks, err
	}
	defer file.Close()

	const errMessage = "Неверный формат данных в файле банков"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		line := strings.TrimSpace(txt)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return banks, errors.New(errMessage)
		}

		name := strings.TrimSpace(parts[0])
		binFrom, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return banks, err
		}
		binTo, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return banks, err
		}

		if binFrom > binTo {
			return banks, errors.New(errMessage)
		}

		banks = append(banks, Bank{
			Name:    name,
			BinFrom: binFrom,
			BinTo:   binTo,
		})
	}

	if err := scanner.Err(); err != nil {
		return banks, err
	}

	return banks, nil
}

// extractBin извлечение BIN из строки номера карты
func extractBIN(cardNumber string) (int, bool) {
	result := 0
	cardNumber = strings.TrimSpace(cardNumber)
	if len(cardNumber) == 0 {
		return result, false
	}

	var builder strings.Builder
	builder.Grow(16)
	for _, s := range cardNumber {
		if s == ' ' {
			continue
		}
		if !unicode.IsDigit(s) {
			return result, false
		}
		builder.WriteRune(s)
	}

	strNum := builder.String()[:6]
	result, err := strconv.Atoi(strNum)
	if err != nil {
		return result, false
	}

	return result, true
}

// identifyBank определение названия банка по его BIN
func identifyBank(bin int, banks []Bank) string {
	result := "Неизвестный банк"

	for _, bank := range banks {
		if bin >= bank.BinFrom && bin <= bank.BinTo {
			result = bank.Name
			break
		}
	}

	return result
}

// validateLuh проверка строки номера карты по алгоритму Луна
func validateLuhn(cardNumber string) bool {
	cardNumber = strings.TrimSpace(cardNumber)
	if len(cardNumber) < 16 {
		return false
	}

	// Шаг 1: Преобразовать входную строку в массив цифр
	var digits []uint8
	for _, s := range cardNumber {
		if s == ' ' {
			continue
		}
		if !unicode.IsDigit(s) {
			return false
		}
		digit, err := strconv.Atoi(string(s))
		if err != nil {
			return false
		}
		digits = append(digits, uint8(digit))
	}

	// Шаг 2: Применить преобразование Луна
	poss := [8]uint8{14, 12, 10, 8, 6, 4, 2, 0}
	for _, pos := range poss {
		p := digits[pos] * 2
		if p > 9 {
			// tens := p / 10
			// ones := p % 10
			// p = tens + ones
			p -= 9
		}
		digits[pos] = p
	}
	var sum uint8
	for _, v := range digits {
		sum += v
	}
	result := (sum % 10) == 0

	return result
}

func main() {
	// banks, err := loadBankData("banks.txt")
	// fmt.Println(banks, err)

	// fmt.Println(identifyBank(400000, banks))

	// bin, ok := extractBIN("9800 1200 7891 8976")
	// if !ok {
	// 	fmt.Println("Номер не определен")
	// 	return
	// }
	// fmt.Println(identifyBank(bin, banks))

	fmt.Println(validateLuhn("4532015112830366"))
	fmt.Println(validateLuhn("1234567890123456"))
	fmt.Println(validateLuhn("9800 1200 7891 8976"))
	fmt.Println(validateLuhn("2202 2062 8242 2422"))
}
