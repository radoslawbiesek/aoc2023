package utils

import (
	"os"
	"strings"
)

func GetContent(path string) (content string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content = string(data)

	return
}

func GetLines(path, separator string) (lines []string) {
	content := GetContent(path)
	lines = strings.Split(content, separator)

	return
}
