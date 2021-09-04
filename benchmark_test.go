package main

import "testing"

func BenchmarkValidatePagePath(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validatePagePath("test")
		validatePagePath("12345")
		validatePagePath("TEST_PAGE")
	}
}

func BenchmarkValidatePagePath2(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validatePagePath2("test")
		validatePagePath2("12345")
		validatePagePath2("TEST_PAGE")
	}
}

func BenchmarkValidatePagePath3(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validatePagePath3("test")
		validatePagePath3("12345")
		validatePagePath3("TEST_PAGE")
	}
}