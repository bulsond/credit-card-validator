package main

import (
	"errors"
	"fmt"
)

type App struct {
	fileBanks string
}

func NewApp(file string) (*App, error) {
	if len(file) == 0 {
		return nil, errors.New("Не указан путь к файлу банков.")
	}
	return &App{fileBanks: file}, nil
}

func (a *App) Run() {
	fmt.Println("Добро пожаловать в программу валидации карт!")
}
