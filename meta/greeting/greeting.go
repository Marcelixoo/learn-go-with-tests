package greeting

import "fmt"

func Hello(name, language string) string {
	return fmt.Sprintf(
		"%s, %s",
		greeting(language),
		name,
	)
}

var greetingByLanguageCode = map[string]string{
	"de": "Hallo",
	"en": "Hello",
	"es": "Hola",
	"fr": "Bonjour",
}

func greeting(language string) string {
	greeting, exists := greetingByLanguageCode[language]

	if !exists {
		return greetingByLanguageCode["en"]
	}

	return greeting
}
