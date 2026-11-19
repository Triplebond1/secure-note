package main

import "testing"

type BenchUser struct {
	Name string
	Age  int
	Data [8]byte
}

var result byte

//go:noinline
func readByValue(user BenchUser) byte {
	return user.Data[0]
}

//go:noinline
func readByPointer(user *BenchUser) byte {
	return user.Data[0]
}

func BenchmarkPassByValue(b *testing.B) {
	user := BenchUser{
		Name: "John",
		Age:  25,
	}

	for i := 0; i < b.N; i++ {
		result = readByValue(user)
	}
}

func BenchmarkPassByPointer(b *testing.B) {
	user := BenchUser{
		Name: "John",
		Age:  25,
	}

	for i := 0; i < b.N; i++ {
		result = readByPointer(&user)
	}
}
