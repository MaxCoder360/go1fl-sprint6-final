package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет это текст или код морзе и переводит
// в противоположный
func Convert(text string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", errors.New("Get an empty string.")
	}

	if isMorseCode(text) {
		// Возвращаем текст, если пришёл код морзе
		return morse.ToText(text), nil
	}

	//  Возвращаем код морзе, если пришёл текст
	return morse.ToMorse(text), nil
}

// isMorseCode проверяет, состоит ли строка только из символов кода Морзе.
func isMorseCode(text string) bool {
	// Проходимся по буквам, и если не находим специальные символы
	// или '.' или '-', то это обычный текст
	for _, r := range text {
		if r == '.' || r == '-' {
			continue
		}

		if r == ' ' || r == '\n' || r == '\r' || r == '\t' {
			continue
		}

		return false
	}

	return true
}
