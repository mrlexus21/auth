// Package config предоставляет функции для загрузки и обработки конфигурационных параметров приложения,
// включая работу с переменными окружения и генерацию конфигураций для подключения к базе данных.
package config

import (
	"github.com/joho/godotenv"
)

// Load загружает конфигурацию по заданному пути.
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
