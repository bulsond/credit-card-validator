package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
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
