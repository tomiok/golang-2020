package main

import (
	"fmt"
	"strings"
)

// string builder

func main() {
	s := join2("hello", ""," my", " ", "name ", "is Tomas")
	fmt.Println(s)
}

func joinWithPlainString(words ...string) string {
	var s string

	for _, w := range words {
		s = s+w
	}

	s = ""

	return s
}

func join(words ...string) string {
	var sb strings.Builder

	for _, w := range words {
		i, _ := sb.WriteString(w)
		fmt.Println(i)
	}
	sb.Reset()

	return sb.String()
}

func join2(words ...string) string {
	var sb strings.Builder

	for _, w := range words {
		fmt.Fprintf(&sb, "%s", w)
	}

	return sb.String()
}