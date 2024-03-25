package main

import "testing"

func Benchmarkfib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}   