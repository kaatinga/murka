package murka

import "testing"

func BenchmarkValidate(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Validate("test")
		Validate("12345")
		Validate("TEST_PAGE", CheckUnderscore)
	}
}

func BenchmarkValidateByRegexp(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validateByRegexp("test")
		validateByRegexp("12345")
		validateByRegexp("TEST_PAGE")
	}
}

func BenchmarkValidateByUnicode(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validateByUnicode("test")
		validateByUnicode("12345")
		validateByUnicode("TEST_PAGE")
	}
}
