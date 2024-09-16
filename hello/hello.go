package hello

import "fmt"

var languages = map[string]language{
	"Spanish": {"Hola, ", "Mundo"},
	"French":  {"Bonjour, ", "tout le monde"},
	"English": {"Hello, ", "World"},
}

func Hello(name string, language string) string {
	var selectedLanguage, languageExists = languages[language]
	if !languageExists {
		selectedLanguage = languages["English"]
	}

	return buildGreeting(selectedLanguage.prefix, name, selectedLanguage.defaultName)
}

func buildGreeting(prefix string, name string, defaultName string) string {
	if name == "" {
		name = defaultName
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}

type language struct {
	prefix      string
	defaultName string
}
