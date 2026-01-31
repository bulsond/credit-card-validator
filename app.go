package main

import (
	"errors"
	"fmt"
)

type App struct {
	fileBanks           string
	minLengthCardNumber int
	maxLengthCardNumber int
}

// NewApp создать экземпляр приложения
func NewApp(file string, minLength, maxLength int) (*App, error) {
	if len(file) == 0 {
		return nil, errors.New("Не указан путь к файлу банков.")
	}
	if minLength < 13 || maxLength > 19 {
		return nil, errors.New("Неверные ограничения на длину номера карты.")
	}

	return &App{
		fileBanks:           file,
		minLengthCardNumber: minLength,
		maxLengthCardNumber: maxLength,
	}, nil
}

// Run запуск приложения
func (a *App) Run() {
	banks, err := loadBankData(a.fileBanks)
	if err != nil {
		panic(err)
	}
	fmt.Println("Добро пожаловать в программу валидации карт!")
	fmt.Println(banks)

	for {
		cardNumber, err := getCardNumber(a.minLengthCardNumber, a.maxLengthCardNumber)
		if err != nil {
			switch err.(type) {
			case ExitRequestError:
				fmt.Println("Программа завершена.")
				return
			case NotValidNumberError:
				fmt.Println("Ошибка:", err)
				continue
			}
		}

		fmt.Printf("Вы ввели: %s\n", cardNumber)
	}
}
