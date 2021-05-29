package main

import "fmt"

const chinese = "Chinese"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const chineseHelloPrefix = "你好, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Chinese"))
}