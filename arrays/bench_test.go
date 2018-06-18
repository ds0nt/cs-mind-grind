package main

import "testing"

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateArray()
	}
}

func BenchmarkSliceNoCapAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateSliceNoCapAppend()
	}
}

func BenchmarkSliceWithCapAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateSliceWithCapAppend()
	}
}

func BenchmarkSliceWithLenCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateSliceWithLenCap()
	}
}

func BenchmarkArrayNoLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateArrayNoLen()
	}
}

func BenchmarkArrayCacheBewm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateArrayCacheBewm()
	}
}

func BenchmarkArrayCacheBewmControl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateArrayCacheBewmControl()
	}
}

func BenchmarkArrayCacheBewm2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateArrayCacheBewm2()
	}
}

func BenchmarkArrayCacheBewm2Control(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateArrayCacheBewm2Control()
	}
}
