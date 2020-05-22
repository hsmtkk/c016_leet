package leet

import (
	"fmt"
	"strings"
)

func Leet(s string) (string, error) {
	if len(s) < 1 || len(s) > 100 {
		return "", fmt.Errorf("too long string")
	}
	var builder strings.Builder
	for _, ch := range s {
		builder.WriteString(transform(string(ch)))
	}
	return builder.String(), nil
}

func transform(ch string) string {
	transformer := map[string]string{
		"A": "4",
		"E": "3",
		"G": "6",
		"I": "1",
		"O": "0",
		"S": "5",
		"Z": "2",
	}
	result, ok := transformer[ch]
	if ok {
		return result
	}
	return ch
}
