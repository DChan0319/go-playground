package hello

import "fmt"
import "strings"

// Hello func returns: string
func Hello(name string, language string) string {

	var prefix string = DeterminePrefix(language)
	var resolvedName string = ResolveName(name, language)

	return prefix + resolvedName
}

// DeterminePrefix func returns: prefix string
func DeterminePrefix(language string) string {
	var prefix string = ""

	if strings.ToLower(language) == "english" {
		prefix = "Hello, "
	} else if strings.ToLower(language) == "spanish" {
		prefix = "Hola, "
	} else if strings.ToLower(language) == "french" {
		prefix = "Bonjour, "
	} else {
		panic("No lanuage prefix for: " + language)
	}

	return prefix
}

// ResolveName func returns: resolvedName string
func ResolveName(name string, language string) string {
	var resolvedName string = name
	if len(strings.TrimSpace(name)) == 0 {
		switch strings.ToLower(language) {
		case "english":
			resolvedName = "World"
		case "spanish":
			resolvedName = "Mundo"
		case "french":
			resolvedName = "Monde"
		}
	}
	return resolvedName
}

func main() {
	fmt.Println(Hello("Darren", "English"))
}
