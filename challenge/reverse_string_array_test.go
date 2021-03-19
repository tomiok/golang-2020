package main

import (
	"testing"
)

func Test_reverseStringArray(t *testing.T) {
	input := "today is the first day of the rest of my life"
	expected := "life my of rest the of day first the is today"
	//  life my of rest the
	result := reverseStringArray(input)

	if result != expected {
		t.Errorf("result: %s", result)
	}
}

func Test_emptySentence(t *testing.T) {
	input, expected := "", ""

	result := reverseStringArray(input)

	if result != expected {
		t.Errorf("result: %s", result)
	}
}

func Test_ZeroIdAlgo(t *testing.T) {
	input := "today is the first day of the rest of my life"
	expected := "life my of rest the of day first the is today"
		    //efil ym fo tser eht fo yad tsrif eht si yadot
	res := reverseZeroIdAlgo(input)

	if res != expected {
		t.Errorf("res %s", res)
	}
}

func Test_reverseAlgo(t *testing.T) {
	input := "today is the first day of the rest of my life"
	expected := "life my of rest the of day first the is today"
	res := reverse(input)
	if res != expected {
		t.Errorf("reverse %s", res)
	}
}

func Benchmark_reverseStringArray(b *testing.B) {
	b.ReportAllocs()
	input := "today is the first day of the rest of my life"
	expected := "life my of rest the of day first the is today"
	result := reverseStringArray(input)

	if result != expected {
		b.Error("")
	}
}
