//following tutorial on https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
package main

const englishPrefix = "Hello, "
const yorubaPrefix = "Enle, "
const hausaPrefix = "Saanu, "
const igboPrefix = "Nnoo, "

//Hello : function to learn unit testing
func Hello(name, language string) string {
	if name == "" {
		name = "Testing"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case "yoruba":
		prefix = yorubaPrefix
	case "hausa":
		prefix = hausaPrefix
	case "igbo":
		prefix = igboPrefix
	default:
		prefix = englishPrefix
	}

	return
}
