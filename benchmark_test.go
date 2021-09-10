package murka

import (
	"strings"
	"testing"
)

//// nolint
//func BenchmarkValidate(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Validate("test")
//		Validate("12345")
//		Validate("TEST_PAGE", CheckUnderscore)
//	}
//}
//
//// nolint
//func BenchmarkValidateByRegexp(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		validateByRegexp("test")
//		validateByRegexp("12345")
//		validateByRegexp("TEST_PAGE")
//	}
//}
//
//// nolint
//func BenchmarkValidateByUnicode(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		validateByUnicode("test")
//		validateByUnicode("12345")
//		validateByUnicode("TEST_PAGE")
//	}
//}

// nolint
func BenchmarkReplace(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Replace("tes:t", '_')
		Replace("12:3:45", '_')
	}
}

// nolint
func BenchmarkStringsReplace(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		strings.ReplaceAll("tes:t", ":", "_")
		strings.ReplaceAll("12:3:45", ":","_")
	}
}