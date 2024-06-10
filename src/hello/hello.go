package main

import "fmt"

const (
	defaultName         = "world"
	englishHelloPrefix  = "Hello, "
	chineseHelloPrefix  = "你好, "
	japaneseHelloPrefix = "こんにちは, "
	chinese             = "chinese"
	japanese            = "japanese"
)

func Hello(language string, name string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", "world"))
}
