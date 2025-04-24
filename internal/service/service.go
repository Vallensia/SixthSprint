package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvMorseString(text string) (string, error) {
	var textOrmorse string
	if text == "" {
		return "", morse.ErrNoEncoding{Text: "empty string"}
	}

	truthCheck := func(r rune) bool {
		return r == '.' || r == '-'
	}
	if strings.ContainsFunc(text, truthCheck) {
		textOrmorse = morse.ToText(text)
	} else {
		textOrmorse = morse.ToMorse(text)
	}

	return textOrmorse, nil
}
