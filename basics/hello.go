package basics

import (
	"fmt"
)

type Language string

const (
	English    Language = "en"
	Portuguese Language = "pt"
	Spanish    Language = "es"
	French     Language = "fr"
)

var languages = map[Language]string{
	English:    "Hello",
	Portuguese: "Olá",
	Spanish:    "Hola",
	French:     "Bonjour",
}

func GetLanguagePrefixSalutation(prefix string) string {
	return languages[Language(prefix)]
}

func SayHello(name string, language string) string {
	return fmt.Sprintf("%s, %s!", GetLanguagePrefixSalutation(language), name)
}
