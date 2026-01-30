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

// Bank
type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

// loadBankData
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

// extractBin
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

// identifyBank
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

func main() {
	banks, err := loadBankData("banks.txt")
	fmt.Println(banks, err)

	fmt.Println(identifyBank(400000, banks))

	bin, ok := extractBIN("9800 1200 7891 8976")
	if !ok {
		fmt.Println("Номер не определен")
		return
	}
	fmt.Println(identifyBank(bin, banks))
}
