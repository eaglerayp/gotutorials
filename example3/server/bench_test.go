package main

import (
	"fmt"
	"testing"
)

func BenchmarkIfLt1(b *testing.B) {
	count := 0
	test := ""
	for n := 0; n < b.N; n++ {
		if len(test) < 1 {
			count++
		}
	}
	fmt.Println("lt1:", count)
}

func BenchmarkIfEq0(b *testing.B) {
	count := 0
	test := ""
	for n := 0; n < b.N; n++ {
		if len(test) == 0 {
			count++
		}
	}
	fmt.Println("Eq0:", count)
}
