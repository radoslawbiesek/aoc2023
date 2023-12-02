package utils

import (
	"fmt"
	"os"
	"strconv"
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

func ParseInt(str string) int {
	parsed, err := strconv.Atoi(strings.TrimSpace(str))

	if err != nil {
		panic(fmt.Sprintf("Could not parse %s", str))
	}

	return parsed
}
