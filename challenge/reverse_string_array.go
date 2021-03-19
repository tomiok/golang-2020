package main

import (
	"strings"
)

func main() {

}

func reverseStringArray(input string) string { // O^n
	values := strings.Fields(input)
	// mantener 2 indices uno con valor 0 y otro con valor del largo del array
	// e ir intercalando los valores a la hora de devolver el array
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}

	return strings.Join(values, " ")
}

func reverseZeroIdAlgo(s string) string {
	r := []rune(s)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
