package main

import "testing"

func BenchmarkSlowEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowEcho()
	}
}

func BenchmarkFastEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastEcho()
	}

}
